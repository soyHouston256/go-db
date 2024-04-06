package main

import (
	"fmt"
	"github.com/soyhouston256/go-db/pkg/product"
	"github.com/soyhouston256/go-db/storage"
	"log"
)

func main() {
	driver := storage.Mysql
	storage.New(driver)
	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("Dao Product Error: %v", err)
	}
	serviceProduct := product.NewService(myStorage)
	ms, err := serviceProduct.GetAll()

	if err != nil {
		log.Fatalf("Service Product Error: %v", err)
	}
	fmt.Printf("Product: %+v\n", ms)

}
