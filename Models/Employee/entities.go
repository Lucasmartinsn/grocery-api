package Employee

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Cpf          int       `json:"cpf"`
	Password     string    `json:"password"`
	Office       string    `json:"office"`
	Active       bool      `json:"active"`
	Admin        bool      `json:"admin"`
	CreationDate time.Time `json:"createon_date"`
}
