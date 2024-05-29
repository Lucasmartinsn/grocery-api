package Employee

import (
	"database/sql"

	"github.com/Lucasmartinsn/grocery-api/Database"
	"github.com/google/uuid"
)

func isUUIDEmpty(u uuid.UUID) bool {
	return u == uuid.Nil
}

func SearchEmployees(id uuid.UUID, status bool) ([]Employee, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if isUUIDEmpty(id) {
		// IF Empty, Get all
		var employee []Employee
		rows, err := conn.Query(`SELECT * FROM t_employee`)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var newemployees Employee
			err = rows.Scan(&newemployees.Id, &newemployees.Name, &newemployees.Cpf, &newemployees.Password, &newemployees.Office,
				&newemployees.Active, &newemployees.CreationDate)
			if err != nil {
				continue
			}
			employee = append(employee, newemployees)
		}
		return employee, err
		// Resquest: http://localhost:5000/api/employee/

	} else if !status && isUUIDEmpty(id) {
		// Get all when tag !empty
		var employee []Employee
		rows, err := conn.Query(`SELECT * FROM t_employee WHERE active=$1`, status)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var newemployees Employee
			err = rows.Scan(&newemployees.Id, &newemployees.Name, &newemployees.Cpf, &newemployees.Password, &newemployees.Office,
				&newemployees.Active, &newemployees.CreationDate)
			if err != nil {
				continue
			}
			employee = append(employee, newemployees)
		}
		return employee, err
		// Resquest: http://localhost:5000/api/employee/?status=true

	} else {
		// Get One
		var employee Employee
		row := conn.QueryRow(`SELECT * FROM t_employee WHERE id=$1`, id)
		err := row.Scan(&employee.Id, &employee.Name, &employee.Cpf, &employee.Password, &employee.Office, &employee.Active, &employee.CreationDate)

		return []Employee{employee}, err
		// Resquest: http://localhost:5000/api/employee/?id=2342
	}
}

func UpdateEmployee(id uuid.UUID, option map[string]bool, employee Employee) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	if option["pass"] {
		// Resquest: http://localhost:5000/api/employee/2342?pass=true
		row, err := conn.Exec(`UPDATE t_employee SET password=$2 WHERE id=$1,`, id, employee.Password)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()

	} else if option["name"] {
		// Resquest: http://localhost:5000/api/employee/2342?name=true
		row, err := conn.Exec(`UPDATE t_employee SET name=$2 WHERE id=$1,`, id, employee.Name)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()

	} else if option["office"] {
		// Resquest: http://localhost:5000/api/employee/2342?office=true
		row, err := conn.Exec(`UPDATE t_employee SET office=$2 WHERE id=$1,`, id, employee.Office)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	} else if option["active"] {
		// Resquest: http://localhost:5000/api/employee/2342?active=true
		row, err := conn.Exec(`UPDATE t_employee SET active=$2 WHERE id=$1,`, id, employee.Active)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	} else if option["admin"] {
		// Resquest: http://localhost:5000/api/employee/2342?admin=true
		row, err := conn.Exec(`UPDATE t_employee SET admin=$2 WHERE id=$1,`, id, employee.Admin)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	}
	return 404, err
}

func DeleteEmployee(id uuid.UUID) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var employee Employee
	err = conn.QueryRow(`SELECT id, admin FROM t_employee WHERE id=$1`, id).Scan(&employee.Id, employee.Admin)
	if err != nil {
		if err == sql.ErrNoRows {
			return 404, err
		}
	}
	
	if employee.Admin {
		// Resquest: http://localhost:5000/api/employee/2342
		response, err := conn.Exec("DELETE FROM t_dominio WHERE id=$1", id)
		if err != nil {
			return 0, err
		}
		return response.RowsAffected()
	}else {
		return 404, err
	}
}

func CreationEmployee(employee Employee) (err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Resquest: http://localhost:5000/api/employee/
	sql := `INSERT INTO t_employee (name, cpf, password, office, admin)
			VALUES ($1,$2,$3,$4,$5)`
	err = conn.QueryRow(sql,&employee.Name, &employee.Cpf, &employee.Password, &employee.Office, &employee.Admin).Err()
	return
}