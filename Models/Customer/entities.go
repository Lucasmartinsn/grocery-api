package Customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id           uuid.UUID `json:"id"`
	Nome         string    `json:"nome"`
	Cpf          int       `json:"cpf"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Office       string    `json:"office"`
	CreationDate time.Time `json:"data_criacao"`
}

type Address struct {
	Id          uuid.UUID `json:"id"`
	Customer_id uuid.UUID `json:"custumer_id"`
	Street      string    `json:"streat"`
	Block       string    `json:"block"`
	Number      int       `json:"number"`
	State       string    `json:"state"`
}

type Credit_card struct {
	Id          uuid.UUID `json:"id"`
	Customer_id uuid.UUID `json:"custumer_id"`
	Number      int       `json:"number"`
	Csv         int       `json:"csvd"`
	NameCard    string    `json:"name_card"`
	Validity    string    `json:"validity"`
}
