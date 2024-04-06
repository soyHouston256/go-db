package storage

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/invoiceheader"
	"github.com/soyhouston256/go-db/pkg/product"
)

const (
	mysqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers (
    		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    		client VARCHAR(100) NOT NULL,
    		created_at TIMESTAMP NOT NULL DEFAULT now(),
    		updated_at TIMESTAMP)`
	mysqlCreateInvoiceHeader = `INSERT INTO invoice_headers (client) VALUES ($1) RETURNING id, created_at`
)

type MysqlInvoiceHeader struct {
	db *sql.DB
}

func (p *MysqlInvoiceHeader) Create(m *product.Model) error {
	//TODO implement me
	panic("implement me")
}

func (p *MysqlInvoiceHeader) GetAll() (product.Models, error) {
	//TODO implement me
	panic("implement me")
}

func (p *MysqlInvoiceHeader) GetByID(id uint) (*product.Model, error) {
	//TODO implement me
	panic("implement me")
}

func (p *MysqlInvoiceHeader) Update(m *product.Model) error {
	//TODO implement me
	panic("implement me")
}

func (p *MysqlInvoiceHeader) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func NewMysqlInvoiceHeader(db *sql.DB) *MysqlInvoiceHeader {
	return &MysqlInvoiceHeader{db}
}

func (p *MysqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Print("Migration invoice header completed")
	return nil
}

func (p *MysqlInvoiceHeader) CreateTx(tx *sql.Tx, model *invoiceheader.Model) error {
	stmt, err := tx.Prepare(mysqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(model.Client).Scan(&model.ID, &model.CreatedAt)
}
