package main

import (
	"github.com/soyhouston256/go-db/pkg/invoice"
	"github.com/soyhouston256/go-db/pkg/invoiceheader"
	"github.com/soyhouston256/go-db/pkg/invoiceitem"
	"github.com/soyhouston256/go-db/storage"
	"log"
)

func main() {
	storage.NewMysqlDB()

	storageheader := storage.NewMysqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewMysqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewMysqlInvoice(
		storage.Pool(),
		storageheader,
		storageItems,
	)
	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "SoyHouston256",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{
				ProductID: 2,
				Quantity:  2,
				Price:     20,
			},
		},
	}
	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("Invoice.Create: %v", err)
	}
}
