package invoiceitem

import (
	"database/sql"
	"time"
)

type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	Quantity        int
	Price           float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	CreateTx(tx *sql.Tx, id uint, models Models) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
