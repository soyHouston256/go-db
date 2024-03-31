package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

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
