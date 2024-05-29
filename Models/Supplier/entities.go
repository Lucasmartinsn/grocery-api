package Supplier

import "github.com/google/uuid"

type Supplier struct {
	Id          uuid.UUID `json:"id"`
	Nome        string    `json:"nome"`
	Cpf         int       `json:"cpf"`
	Cnpj        int       `json:"cnpj"`
	CompanyName string    `json:"company_name"`
}

type Product struct {
	Id       uuid.UUID `json:"id"`
	Nome     string    `json:"nome"`
	Batch_id uuid.UUID `json:"supplier_id"`
	Price    float64
	Validity string
}

type Batch struct {
	Id               uuid.UUID `json:"id"`
	Supplier_id      uuid.UUID `json:"supplier_id"`
	Volume           int       `json:"volume"`
	ValidityProduct1 string    `json:"validity_product1"`
	ValidityProduct2 string    `json:"validity_product2"`
	DeliveryDate     string    `json:"delivery_date"`
}
