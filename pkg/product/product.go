package product

import "time"

type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	Create(m *Model) error
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
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}
