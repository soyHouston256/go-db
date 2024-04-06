package storage

import (
	"database/sql"
	"fmt"
	"github.com/soyhouston256/go-db/pkg/product"
)

const (
	mysqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name TEXT,
    observation TEXT,
    price FLOAT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
    )`
	mysqlCreateProduct  = `INSERT INTO products(name, observation, price, created_at) VALUES(?, ?, ?, ?)`
	mysqlGetAllProduct  = `SELECT id, name, observation, price, created_at, updated_at FROM products`
	mysqlGetProductByID = mysqlGetAllProduct + " WHERE id = ?"
)

type MysqlProduct struct {
	db *sql.DB
}

func (p *MysqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mysqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observation),
		m.Price,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	if err != nil {
		return err
	}
	m.ID = uint(id)
	fmt.Printf("Product created with ID: %d\n", m.ID)
	return nil
}

func (p *MysqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(mysqlGetAllProduct)
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

func (p *MysqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(mysqlGetProductByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return ScanRowProduct(stmt.QueryRow(id))
}

func (p *MysqlProduct) Update(m *product.Model) error {
	//TODO implement me
	panic("implement me")
}

func (p *MysqlProduct) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func NewMysqlProduct(db *sql.DB) *MysqlProduct {
	return &MysqlProduct{db}
}
func (p *MysqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateProduct)
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
