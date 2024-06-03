package Supplier

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	Id              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Cnpj            int       `json:"cnpj"`
	Contract_number int       `json:"contract_number"`
	CompanyName     string    `json:"company_name"`
	Status          bool      `json:"status"`
}

type Product struct {
	Id          uuid.UUID `json:"id"`
	Batch_id    uuid.UUID `json:"batch_id"`
	Supplier_id uuid.UUID `json:"supplier_id"`
	Name        string    `json:"name"`
	Volume      int       `json:"volume"`
	Unit_price  float64   `json:"unit_price"`
	Validity    time.Time `json:"validity"`
}

type Batch struct {
	Id            uuid.UUID `json:"id"`
	Supplier_id   uuid.UUID `json:"supplier_id"`
	Volume        int       `json:"volume"`
	Price         float64   `json:"price"`
	Purchase_date time.Time `json:"purchase_date"`
	DeliveryDate  time.Time `json:"delivery_date"`
}

type S_product struct {
	Supplier Supplier
	Products []Product
}
type S_batch struct {
	Supplier Supplier
	Batchs   []Batch
}
