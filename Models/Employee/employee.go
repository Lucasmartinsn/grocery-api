package Employee

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/Lucasmartinsn/grocery-api/Database"
	Services "github.com/Lucasmartinsn/grocery-api/Services/EncryptionPass"
	"github.com/google/uuid"
)

func isUUIDEmpty(u uuid.UUID) bool {
	// uuid.Nil == Null
	return u == uuid.Nil
}

func SearchEmployees(i, s string) ([]Employee, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	id, _ := uuid.Parse(i)
	status, _ := strconv.ParseBool(s)
	if i == "" && s == "" {
		// Resquest: http://localhost:5000/api/employee/
		var employee []Employee
		rows, err := conn.Query(`SELECT id, name, cpf, office, active, admin, creation_date FROM t_employee`)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var newemployees Employee
			err = rows.Scan(&newemployees.Id, &newemployees.Name, &newemployees.Cpf, &newemployees.Office,
				&newemployees.Active, &newemployees.Admin, &newemployees.CreationDate)
			if err != nil {
				continue
			}
			employee = append(employee, newemployees)
		}
		return employee, err

	} else if !isUUIDEmpty(id) && s == "" {
		// Get One
		// Resquest: http://localhost:5000/api/employee/?id=2342
		var employee Employee
		row := conn.QueryRow(`SELECT id, name, cpf, office, active, admin, creation_date FROM t_employee WHERE id=$1`, id)
		err := row.Scan(&employee.Id, &employee.Name, &employee.Cpf, &employee.Office, &employee.Active, &employee.Admin, &employee.CreationDate)

		return []Employee{employee}, err

	} else if i == "" && s != "" {
		// Get all when tag !empty
		// Resquest: http://localhost:5000/api/employee/?status=true
		var employee []Employee
		rows, err := conn.Query(`SELECT id, name, cpf, office, active, admin, creation_date FROM t_employee WHERE active=$1`, status)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var newemployees Employee
			err = rows.Scan(&newemployees.Id, &newemployees.Name, &newemployees.Cpf, &newemployees.Office,
				&newemployees.Active, &newemployees.Admin, &newemployees.CreationDate)
			if err != nil {
				continue
			}
			employee = append(employee, newemployees)
		}
		return employee, err

	} else if i != "" && s != "" {
		// 	// Get One
		// Resquest: http://localhost:5000/api/employee/?id=2342&status=true
		var employee Employee
		row := conn.QueryRow(`SELECT id, name, cpf, office, active, admin, creation_date FROM t_employee WHERE id=$1 and active=$2`, id, status)
		err := row.Scan(&employee.Id, &employee.Name, &employee.Cpf, &employee.Office, &employee.Active, &employee.Admin, &employee.CreationDate)

		return []Employee{employee}, err
	}

	return []Employee{}, errors.New("no conditions met")
}

func UpdateEmployee(id uuid.UUID, option map[string]bool, employee Employee) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	if option["all"] {
		// Resquest: http://localhost:5000/api/employee/2342?all=true
		password, _ := Services.HashPassword(employee.Password)
		row, err := conn.Exec(
			`UPDATE t_employee SET name=$2, password=$3, office=$4, active=$5, admin=$6 WHERE id=$1`, id,
			employee.Name, password, employee.Office, employee.Active, employee.Admin)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()

	} else if option["pass"] {
		// Resquest: http://localhost:5000/api/employee/2342?pass=true
		password, _ := Services.HashPassword(employee.Password)
		row, err := conn.Exec(`UPDATE t_employee SET password=$2 WHERE id=$1`, id, password)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()

	} else if option["name"] {
		// Resquest: http://localhost:5000/api/employee/2342?name=true
		row, err := conn.Exec(`UPDATE t_employee SET name=$2 WHERE id=$1`, id, employee.Name)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()

	} else if option["office"] {
		// Resquest: http://localhost:5000/api/employee/2342?office=true
		row, err := conn.Exec(`UPDATE t_employee SET office=$2 WHERE id=$1`, id, employee.Office)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	} else if option["active"] {
		// Resquest: http://localhost:5000/api/employee/2342?active=true
		row, err := conn.Exec(`UPDATE t_employee SET active=$2 WHERE id=$1`, id, employee.Active)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	} else if option["admin"] {
		// Resquest: http://localhost:5000/api/employee/2342?admin=true
		row, err := conn.Exec(`UPDATE t_employee SET admin=$2 WHERE id=$1`, id, employee.Admin)
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
	err = conn.QueryRow(`SELECT id, admin FROM t_employee WHERE id=$1`, id).Scan(&employee.Id, &employee.Admin)
	if err != nil {
		if err == sql.ErrNoRows {
			return 404, err
		}
	}

	if employee.Admin {
		// Resquest: http://localhost:5000/api/employee/2342
		response, err := conn.Exec("DELETE FROM t_employee WHERE id=$1", id)
		if err != nil {
			return 0, err
		}
		return response.RowsAffected()
	} else {
		return 404, err
	}
}

func CreationEmployee(employee Employee) (err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	password, _ := Services.HashPassword(employee.Password)
	// Resquest: http://localhost:5000/api/employee/
	sql := `INSERT INTO t_employee (name, cpf, password, office, active, admin)
			VALUES ($1,$2,$3,$4,$5,$6)`
	err = conn.QueryRow(sql, &employee.Name, &employee.Cpf, password, &employee.Office, &employee.Active, &employee.Admin).Err()
	return
}

func Validate(cpf int, senha string) (employee Employee, err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT id, password, active FROM t_employee WHERE cpf=$1`
	err = conn.QueryRow(query, cpf).Scan(&employee.Id, &employee.Password, &employee.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("user does not exist")
		}
		return
	}
	err = Services.CheckPassword(employee.Password, senha)
	if err != nil {
		err = errors.New("incorrect username or password")
		return
	} else if !employee.Active {
		err = errors.New("inactive user")
		return
	}
	return
}
