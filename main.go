package main

import (
	"github.com/soyhouston256/go-db/pkg/product"
	"github.com/soyhouston256/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	m := &product.Model{
		Name:  "Course of Go",
		Price: 120.50,
	}
	if err := serviceProduct.Create(m); err != nil {
		panic(err)
	}

}
