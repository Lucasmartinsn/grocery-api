package Customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Contact      int       `json:"contact"`
	Cpf          int       `json:"cpf"`
	Password     string    `json:"password"`
	CreationDate time.Time `json:"creation_date"`
}

type Address struct {
	Id          uuid.UUID `json:"id"`
	Customer_id uuid.UUID `json:"customer_id"`
	Street      string    `json:"street"`
	Block       string    `json:"block"`
	Number      int       `json:"number"`
	State       string    `json:"state"`
}

type Credit_card struct {
	Id          uuid.UUID `json:"id"`
	Customer_id uuid.UUID `json:"customer_id"`
	Number      int       `json:"number"`
	Csv         int       `json:"csv"`
	NameCard    string    `json:"name_card"`
	Validity    string    `json:"validity"`
}

type C_Address struct {
	Customer Customer
	Address  []Address
}

type C_Card struct {
	Customer Customer
	Cards    []Credit_card
}
