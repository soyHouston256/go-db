package main

import (
	"fmt"
	"github.com/soyhouston256/go-db/pkg/invoiceheader"
	"github.com/soyhouston256/go-db/pkg/invoiceitem"
	"github.com/soyhouston256/go-db/pkg/product"
	"github.com/soyhouston256/go-db/storage"
)

func main() {
	storage.NewMysqlDB()

	storageProduct := storage.NewMysqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	serviceProduct.Migrate()

	storageHeader := storage.NewMysqlInvoiceHeader(storage.Pool())
	serviceHeader := invoiceheader.NewService(storageHeader)
	serviceHeader.Migrate()

	storageInvoiceItem := storage.NewMysqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	serviceInvoiceItem.Migrate()

	fmt.Printf("Migration product completed")
}
