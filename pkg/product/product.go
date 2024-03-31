package product

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrNotFound = errors.New("item not found")
)

type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s",
		m.ID, m.Name, m.Observation, m.Price,
		m.CreatedAt.Format("2004-01-01"),
		m.UpdatedAt.Format("2004-01-01"))
}

type Models []*Model

func (m Models) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-20s | %5s | %10s | %10s\n",
		"id", "name", "observation", "price", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

type Storage interface {
	Migrate() error
	Create(m *Model) error
	GetAll() (Models, error)
	GetByID(id uint) (*Model, error)
	Update(m *Model) error
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

func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrNotFound
	}
	m.UpdatedAt = time.Now()
	return s.storage.Update(m)
}
