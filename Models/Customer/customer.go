package Customer

import (
	"database/sql"
	"encoding/json"
	"errors"

	confDB "github.com/Lucasmartinsn/grocery-api/Configs/confEnv"
	"github.com/Lucasmartinsn/grocery-api/Database"
	Type "github.com/Lucasmartinsn/grocery-api/Models/Employee"
	Services "github.com/Lucasmartinsn/grocery-api/Services/EncryptionPass"
	"github.com/Lucasmartinsn/grocery-api/Services/EncryptionResponse"
	"github.com/google/uuid"
)

func convert(v any) (string, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}
func ValidateEmployee(cpf int, senha string) (employee Type.Employee, err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT id, password, active FROM t_employee WHERE cpf=$1`
	err = conn.QueryRow(query, cpf).Scan(&employee.Id, &employee.Password, &employee.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("incorrect username or password")
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
func ValidateCustomer(cpf int, senha string) (customer Customer, err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT id, password FROM t_customer WHERE cpf=$1`
	err = conn.QueryRow(query, cpf).Scan(&customer.Id, &customer.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("incorrect username or password")
		}
		return
	}
	match, err := Services.VerifyHash(senha, customer.Password)
	if err != nil || !match {
		err = errors.New("incorrect username or password")
		return
	}
	return
}
func CreationCustomer(customer Customer) (id uuid.UUID, err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	// Resquest: http://localhost:5000/api/customer/?customer=true
	Password, _ := Services.GenerateHash(customer.Password)
	sql := `INSERT INTO t_customer (name, email, contact, cpf, password)
			VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err = conn.QueryRow(sql, &customer.Name, &customer.Email, &customer.Contact, &customer.Cpf, Password).Scan(&id)
	return
}
func CreationAddress(address Address) (err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Resquest: http://localhost:5000/api/customer/?address=true
	sql := `INSERT INTO t_address (customer_id, street, block, number, state)
			VALUES ($1,$2,$3,$4,$5)`
	err = conn.QueryRow(sql, &address.Customer_id, &address.Street, &address.Block, &address.Number, &address.State).Err()
	return
}
func CreationCard(card Credit_card) (err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/customer/?card=true
	sql := `INSERT INTO t_credit_card (customer_id, number, csv, name_card, validity)
			VALUES ($1,$2,$3,$4,$5)`
	err = conn.QueryRow(sql, &card.Customer_id, &card.Number, &card.Csv, &card.NameCard, &card.Validity).Err()
	return
}
func SearchCustomer(ids string) (string, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/customer/?id=21331
	if id, _ := uuid.Parse(ids); id != uuid.Nil {
		var customer Customer
		row := conn.QueryRow(`SELECT id, name, email, contact, cpf, creation_date FROM t_customer WHERE id=$1`, id)
		err = row.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Contact, &customer.Cpf, &customer.CreationDate)
		if err != nil {
			return "", err
		}
		data, err := convert(customer)
		if err != nil {
			return "", err
		}
		return EncryptionResponse.EncryptData(data, []byte(confDB.Variable()))

	} else {
		// Resquest: http://localhost:5000/api/customer/
		var customer []Customer
		rows, err := conn.Query(`SELECT id, name, email, contact, cpf, creation_date FROM t_customer`)
		if err != nil {
			return "", err
		}

		for rows.Next() {
			var newcustomer Customer
			err = rows.Scan(&newcustomer.Id, &newcustomer.Name, &newcustomer.Email, &newcustomer.Contact,
				&newcustomer.Cpf, &newcustomer.CreationDate)
			if err != nil {
				continue
			}
			customer = append(customer, newcustomer)
		}
		data, err := convert(customer)
		if err != nil {
			return "", err
		}
		return EncryptionResponse.EncryptData(data, []byte(confDB.Variable()))
	}
}
func SearchCustomer_address(id uuid.UUID) (string, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Resquest: http://localhost:5000/api/customer/address/1233
	var customer Customer
	row := conn.QueryRow(`SELECT id, name, email, contact, cpf, creation_date FROM t_customer WHERE id=$1`, id)
	err = row.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Contact, &customer.Cpf, &customer.CreationDate)
	if err != nil {
		return "", err
	}
	var address []Address
	rows, err := conn.Query(`SELECT * FROM t_address WHERE customer_id=$1`, id)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		var newaddress Address
		err = rows.Scan(&newaddress.Id, &newaddress.Customer_id, &newaddress.Street, &newaddress.Block, &newaddress.Number, &newaddress.State)
		if err != nil {
			continue
		}
		address = append(address, newaddress)
	}
	s_address := C_Address{
		Customer: customer,
		Address:  address,
	}
	data, err := convert(s_address)
	if err != nil {
		return "", err
	}
	return EncryptionResponse.EncryptData(data, []byte(confDB.Variable()))
}
func SearchCustomer_card(id uuid.UUID) (string, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Resquest: http://localhost:5000/api/customer/card/1233
	var customer Customer
	row := conn.QueryRow(`SELECT id, name, email, contact, cpf, creation_date FROM t_customer WHERE id=$1`, id)
	err = row.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Contact, &customer.Cpf, &customer.CreationDate)
	if err != nil {
		return "", err
	}
	var card []Credit_card
	rows, err := conn.Query(`SELECT * FROM t_credit_card WHERE customer_id=$1`, id)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		var newcard Credit_card
		err = rows.Scan(&newcard.Id, &newcard.Customer_id, &newcard.Number, &newcard.Csv, &newcard.NameCard, &newcard.Validity)
		if err != nil {
			continue
		}
		card = append(card, newcard)
	}
	s_card := C_Card{
		Customer: customer,
		Cards:    card,
	}
	data, err := convert(s_card)
	if err != nil {
		return "", err
	}
	return EncryptionResponse.EncryptData(data, []byte(confDB.Variable()))
}
func UpdatedCustomer(id uuid.UUID, customer Customer, p bool) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/customer/2342?customer=true&pass=true
	if p {
		pass, _ := Services.GenerateHash(customer.Password)
		row, err := conn.Exec(
			`UPDATE t_customer SET password=$2 WHERE id=$1`, id, pass)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	} else {
		// Resquest: http://localhost:5000/api/customer/2342?customer=true
		row, err := conn.Exec(
			`UPDATE t_customer SET name=$2, email=$3, contact=$4, cpf=$5 WHERE id=$1`, id,
			customer.Name, customer.Email, customer.Contact, customer.Cpf)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	}
}
func UpdatedCustomer_address(id uuid.UUID, address Address) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/customer/2342?address=true
	row, err := conn.Exec(
		`UPDATE t_address SET street=$2, block=$3, number=$4, state=$5 WHERE id=$1`, id, address.Street, address.Block, address.Number, address.State)
	if err != nil {
		return 0, err
	}
	return row.RowsAffected()
}
func UpdatedCustomer_card(id uuid.UUID, card Credit_card) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/customer/2342?card=true
	row, err := conn.Exec(
		`UPDATE t_credit_card SET number=$2, csv=$3, name_card=$4, validity=$5 WHERE id=$1`, id,
		card.Number, card.Csv, card.NameCard, card.Validity)
	if err != nil {
		return 0, err
	}
	return row.RowsAffected()
}
func DeleteCustumer(id uuid.UUID) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/customer/2342
	response, err := conn.Exec("DELETE FROM t_customer WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return response.RowsAffected()
}
