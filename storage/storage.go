package storage

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/product"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

type Driver string

const (
	Mysql    Driver = "mysql"
	Postgres Driver = "postgres"
)

func New(d Driver) {
	switch d {
	case Mysql:
		newMysqlDB()
	case Postgres:
		newPostgresDB()
	default:
		newMysqlDB()
	}
}

func newPostgresDB() {
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

func newMysqlDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1)/golang?parseTime=true")

		if err != nil {
			log.Fatalf("cant open db connection: %v", err)
			panic(err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("cant do ping: %v", err)
		}
		fmt.Println("connected to mysql db")
	})
}

func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	var ns sql.NullString
	if s == "" {
		ns.Valid = false
		return ns
	}
	ns.String = s
	ns.Valid = true
	return ns
}

func timeToNull(t time.Time) sql.NullTime {
	var nt sql.NullTime
	if t.IsZero() {
		nt.Valid = false
		return nt
	}
	nt.Time = t
	nt.Valid = true
	return nt
}

func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Mysql:
		return newMysqlProduct(db), nil
	case Postgres:
		return newPsqlProduct(db), nil
	default:
		return newMysqlProduct(db), nil
	}
}
