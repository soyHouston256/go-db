package storage

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

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

	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"

	psqlUpdateProduct = `UPDATE products SET name=$1, observation=$2, price=$3, updated_at=$4 WHERE id=$5`

	psqlDeleteProduct = `DELETE FROM products WHERE id=$1`
)

type psqlProduct struct {
	db *sql.DB
}

func newPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}
func (p *psqlProduct) Migrate() error {
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

func (p *psqlProduct) Create(m *product.Model) error {
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

func (p *psqlProduct) GetAll() (product.Models, error) {
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
		m, err := ScanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

func (p *psqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return ScanRowProduct(stmt.QueryRow(id))
}

func (p *psqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observation),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("not exist the product with ID: %d", m.ID)
	}
	fmt.Print("Product updated")
	return nil
}

func (p *psqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	fmt.Print("Product deleted")
	return nil
}
func ScanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull)
	if err != nil {
		return &product.Model{}, err
	}
	m.Observation = observationNull.String
	m.UpdatedAt = updatedAtNull.Time
	return m, nil
}
