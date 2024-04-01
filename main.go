package main

import (
	"github.com/soyhouston256/go-db/pkg/invoice"
	"github.com/soyhouston256/go-db/pkg/invoiceheader"
	"github.com/soyhouston256/go-db/pkg/invoiceitem"
	"github.com/soyhouston256/go-db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()
	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(storage.Pool(), storageHeader, storageItems)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "HOUSTON RAMIREZ",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{
				ProductID: 2,
				Price:     25,
				Quantity:  1,
			},
			&invoiceitem.Model{
				ProductID: 2,
				Price:     27,
				Quantity:  2,
			},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

}
