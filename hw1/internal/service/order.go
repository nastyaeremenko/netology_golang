package service

import (
	"fmt"

	"example/solid/internal/notifier"
	"example/solid/internal/repository"
	"example/solid/internal/repository/model"
)

// OrderService - бизнес-логика, зависит только от интерфейсов
type OrderService struct {
	repo     repository.RepositoryWriter
	notifier notifier.Notifier
}

// в параметрах - интерфейсы, а не конкретные реализации
func New(repo repository.RepositoryWriter, notifier notifier.Notifier) *OrderService {
	return &OrderService{repo: repo, notifier: notifier}
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	order := model.Order{
		Customer: customer,
		Products: fmt.Sprintf("%v", products),
		Total:    total,
		Status:   "pending",
	}

	err := s.repo.SaveOrder(order)
	if err != nil {
		return err
	}

	return s.notifier.Send(customer)
}
