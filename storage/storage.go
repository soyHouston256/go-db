package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error

		db, err = sql.Open("postgres", "postgres://user:admin@localhost:5432/user?sslmode=disable")

		if err != nil {
			log.Fatalf("cant open db connection: %v", err)
			panic(err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("cant do ping: %v", err)
		}
		fmt.Println("connected to db")
	})
}

func Pool() *sql.DB {
	return db
}
