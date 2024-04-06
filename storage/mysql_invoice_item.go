package storage

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/invoiceitem"
)

const (
	mysqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items (
    		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    		invoice_header_id INT NOT NULL,
    		product_id INT NOT NULL,
    		quantity INT,
    		price FLOAT,
    		created_at TIMESTAMP NOT NULL DEFAULT now(),
    		updated_at TIMESTAMP,
    		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
    		)`
	mysqlCreateInvoiceItem = `INSERT INTO invoice_items (invoice_header_id, product_id, quantity, price) VALUES (?, ?, ?, ?)`
)

type MysqlInvoiceItem struct {
	db *sql.DB
}

func NewMysqlInvoiceItem(db *sql.DB) *MysqlInvoiceItem {
	return &MysqlInvoiceItem{db}
}

func (p *MysqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Print("Migration invoice item completed")
	return nil
}

func (p *MysqlInvoiceItem) CreateTx(tx *sql.Tx, invoiceHeaderID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(mysqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		result, err := stmt.Exec(invoiceHeaderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		item.ID = uint(id)
	}
	return nil
}
