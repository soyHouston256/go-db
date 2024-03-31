package main

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/product"
	"github.com/soyhouston256/go-db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	ms, err := serviceProduct.GetByID(0)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case err != nil:
		log.Fatalf("product.GetByID: %v", err)
	default:
		fmt.Println(ms)
	}

}
