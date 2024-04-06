package storage

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/invoiceheader"
)

const (
	mysqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers (
    		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    		client VARCHAR(100) NOT NULL,
    		created_at TIMESTAMP NOT NULL DEFAULT now(),
    		updated_at TIMESTAMP)`
	mysqlCreateInvoiceHeader = `INSERT INTO invoice_headers (client) VALUES(?)`
)

type MysqlInvoiceHeader struct {
	db *sql.DB
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

	result, err := stmt.Exec(model.Client)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	model.ID = uint(id)
	fmt.Printf("Invoice header created with ID: %d\n", model.ID)
	return nil
}
