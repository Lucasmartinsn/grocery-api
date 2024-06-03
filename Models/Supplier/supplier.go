package Supplier

import (
	"github.com/Lucasmartinsn/grocery-api/Database"
	"github.com/google/uuid"
)

func CreationSupplier(supplier Supplier) (err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/supplier/?supplier=true
	sql := `INSERT INTO t_supplier (name, cnpj, contract_number, company_name, status)
			VALUES ($1,$2,$3,$4,$5)`
	err = conn.QueryRow(sql, &supplier.Name, &supplier.Cnpj, &supplier.Contract_number, &supplier.CompanyName, &supplier.Status).Err()
	return
}
func CreationProduct(product Product) (err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Resquest: http://localhost:5000/api/supplier/?product=true
	sql := `INSERT INTO t_product (batch_id, supplier_id, name, volume, unit_price, validity)
			VALUES ($1,$2,$3,$4,$5,$6)`
	err = conn.QueryRow(sql, &product.Batch_id, &product.Supplier_id, &product.Name, &product.Volume, &product.Unit_price, &product.Validity).Err()
	return
}
func CreationBatch(batch Batch) (err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/supplier/?batch=true
	sql := `INSERT INTO t_batch (supplier_id, volume, price, delivery_date)
			VALUES ($1,$2,$3,$4)`
	err = conn.QueryRow(sql, &batch.Supplier_id, &batch.Volume, &batch.Price, &batch.DeliveryDate).Err()
	return
}
func SearchSupplier() (supplier []Supplier, err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/supplier/
	rows, err := conn.Query(`SELECT * FROM t_supplier`)
	if err != nil {
		return
	}
	for rows.Next() {
		var newsupplier Supplier
		err = rows.Scan(&newsupplier.Id, &newsupplier.Name, &newsupplier.Cnpj, &newsupplier.Contract_number,
			&newsupplier.CompanyName, &newsupplier.Status)
		if err != nil {
			continue
		}
		supplier = append(supplier, newsupplier)
	}
	return
}
func SearchSupplier_product(id uuid.UUID) (s_product S_product, err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// Resquest: http://localhost:5000/api/supplier/product/1233
	var supplier Supplier
	row := conn.QueryRow(`SELECT * FROM t_supplier WHERE id=$1`, id)
	err = row.Scan(&supplier.Id, &supplier.Name, &supplier.Cnpj, &supplier.Contract_number, &supplier.CompanyName, &supplier.Status)
	if err != nil {
		return
	}
	var product []Product
	rows, err := conn.Query(`SELECT * FROM t_product WHERE supplier_id=$1`, id)
	if err != nil {
		return
	}
	for rows.Next() {
		var newproduct Product
		err = rows.Scan(&newproduct.Id, &newproduct.Batch_id, &newproduct.Supplier_id, &newproduct.Name, &newproduct.Volume, &newproduct.Unit_price, &newproduct.Validity)
		if err != nil {
			continue
		}
		product = append(product, newproduct)
	}
	s_product = S_product{
		Supplier: supplier,
		Products: product,
	}
	return s_product, err
}
func SearchSupplier_bacth(id uuid.UUID) (s_batch S_batch, err error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// Resquest: http://localhost:5000/api/supplier/batch/1233
	var supplier Supplier
	row := conn.QueryRow(`SELECT * FROM t_supplier WHERE id=$1`, id)
	err = row.Scan(&supplier.Id, &supplier.Name, &supplier.Cnpj, &supplier.Contract_number, &supplier.CompanyName, &supplier.Status)
	if err != nil {
		return
	}
	var batch []Batch
	rows, err := conn.Query(`SELECT * FROM t_batch WHERE supplier_id=$1`, id)
	if err != nil {
		return
	}
	for rows.Next() {
		var newbatch Batch
		err = rows.Scan(&newbatch.Id, &newbatch.Supplier_id, &newbatch.Volume, &newbatch.Price, &newbatch.Purchase_date, &newbatch.DeliveryDate)
		if err != nil {
			continue
		}
		batch = append(batch, newbatch)
	}
	s_batch = S_batch{
		Supplier: supplier,
		Batchs:   batch,
	}
	return s_batch, err
}
func UpdatedSupplier(id uuid.UUID, supplier Supplier) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/supplier/2342?supplier=true
	row, err := conn.Exec(
		`UPDATE t_supplier SET name=$2, cnpj=$3, contract_number=$4, company_name=$5, status=$6 WHERE id=$1`, id,
		supplier.Name, supplier.Cnpj, supplier.Contract_number, supplier.CompanyName, supplier.Status)
	if err != nil {
		return 0, err
	}
	return row.RowsAffected()
}
func UpdatedProduct(id uuid.UUID, product Product, v bool) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	if v {
		// Resquest: http://localhost:5000/api/supplier/2342?product=true&volume=true
		var valor int
		prod := conn.QueryRow(`SELECT volume FROM t_product WHERE id=$1`, id)
		err = prod.Scan(&valor)
		if err != nil {
			return 500, err
		}
		real := valor - product.Volume
		row, err := conn.Exec(
			`UPDATE t_product SET volume=$2 WHERE id=$1`, id, real)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()

	} else {
		// Resquest: http://localhost:5000/api/supplier/2342?product=true
		row, err := conn.Exec(
			`UPDATE t_product SET name=$2, volume=$3, unit_price=$4 WHERE id=$1`, id, product.Name, product.Volume, product.Unit_price)
		if err != nil {
			return 0, err
		}
		return row.RowsAffected()
	}
}
func UpdatedBatch(id uuid.UUID, batch Batch) (int64, error) {
	conn, err := Database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	// Resquest: http://localhost:5000/api/supplier/2342?batch=true
	row, err := conn.Exec(
		`UPDATE t_batch SET volume=$2, price=$3, delivery_date=$4 WHERE id=$1`, id,
		batch.Volume, batch.Price, batch.DeliveryDate)
	if err != nil {
		return 0, err
	}
	return row.RowsAffected()
}
