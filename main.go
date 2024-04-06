package main

import (
	"fmt"
	"github.com/soyhouston256/go-db/pkg/product"
	"github.com/soyhouston256/go-db/storage"
)

func main() {
	storage.NewMysqlDB()

	storageProduct := storage.NewMysqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	m := &product.Model{
		Name:        `Product 1`,
		Observation: `This is the first product`,
		Price:       100,
	}
	if err := serviceProduct.Create(m); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Migration product completed")
}
