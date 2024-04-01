package storage

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/invoice"
	"github.com/soyhouston256/go-db/pkg/invoiceheader"
	"github.com/soyhouston256/go-db/pkg/invoiceitem"
)

type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

func NewPsqlInvoice(db *sql.DB, storageHeader invoiceheader.Storage, storageItems invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{db, storageHeader, storageItems}
}

func (p *PsqlInvoice) Create(model *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, model.Header); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Header created with ID: %d\n", model.Header.ID)

	if err := p.storageItems.CreateTx(tx, model.Header.ID, model.Items); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Items created:  %d\n", len(model.Items))
	return tx.Commit()
}
