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

	ms, err := serviceProduct.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", ms)
}
