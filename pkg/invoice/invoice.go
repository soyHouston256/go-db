package invoice

import (
	"github.com/soyhouston256/go-db/pkg/invoiceheader"
	"github.com/soyhouston256/go-db/pkg/invoiceitem"
)

type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

type Storage interface {
	Create(model *Model) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Create(model *Model) error {
	return s.storage.Create(model)
}
