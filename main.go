package main

import (
	"github.com/soyhouston256/go-db/pkg/product"
	"github.com/soyhouston256/go-db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	m := &product.Model{
		ID:    5,
		Name:  "test",
		Price: 100,
	}
	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.GetByID: %v", err)
	}
}
