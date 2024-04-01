package storage

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/invoiceitem"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items (
    		id SERIAL NOT NULL,
    		invoice_header_id INT NOT NULL,
    		product_id INT NOT NULL,
    		quantity INT,
    		price FLOAT,
    		created_at TIMESTAMP NOT NULL DEFAULT now(),
    		updated_at TIMESTAMP,
    		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
    		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
    		)`
	psqlCreateInvoiceItem = `INSERT INTO invoice_items (invoice_header_id, product_id, quantity, price) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
)

type PsqlInvoiceItem struct {
	db *sql.DB
}

func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
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

func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, invoiceHeaderID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		err = stmt.QueryRow(invoiceHeaderID, item.ProductID, item.Quantity, item.Price).Scan(&item.ID, &item.CreatedAt)
		if err != nil {
			return err
		}
	}
	return nil
}
