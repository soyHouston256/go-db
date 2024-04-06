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
)

type MysqlProduct struct {
	db *sql.DB
}

func (p *MysqlProduct) Create(m *product.Model) error {
	//TODO implement me
	panic("implement me")
}

func (p *MysqlProduct) GetAll() (product.Models, error) {
	//TODO implement me
	panic("implement me")
}

func (p *MysqlProduct) GetByID(id uint) (*product.Model, error) {
	//TODO implement me
	panic("implement me")
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
