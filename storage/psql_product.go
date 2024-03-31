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

	psqlGetAllProduct = `SELECT id, name, observation, price, created_at, updated_at FROM products`
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

func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := make(product.Models, 0)
	for rows.Next() {
		m := &product.Model{}
		observationNull := sql.NullString{}
		updatedAtNull := sql.NullTime{}
		err := rows.Scan(&m.ID, &m.Name, &observationNull, &m.Price, &m.CreatedAt, &updatedAtNull)
		if err != nil {
			return nil, err
		}

		m.Observation = observationNull.String
		m.UpdatedAt = updatedAtNull.Time

		ms = append(ms, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}
