package storage

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products (
    id SERIAL NOT NULL,
    name TEXT,
    observation TEXT,
    price FLOAT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT products_id_pk PRIMARY KEY (id)
    )`

	pqsCreateProduct = `INSERT INTO products (name, observation, price, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
)

type PsqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Print("Migration product completed")
	return nil
}

func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(pqsCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(m.Name, stringToNull(m.Observation), m.Price, m.CreatedAt).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Print("Product created")
	return nil
}
