package main

import (
	"github.com/soyhouston256/go-db/pkg/product"
	"github.com/soyhouston256/go-db/storage"
	"log"
)

func main() {
	storage.NewMysqlDB()

	storageProduct := storage.NewMysqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:          2,
		Name:        "course of Go",
		Observation: "course of Go beginner to master",
		Price:       120,
	}

	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}
