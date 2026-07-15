package service

import (
	"testing"

	"example/solid/internal/repository/model"
)

// MockWriter - мок репозитория, ничего не пишет в базу
type MockWriter struct {
	saved []model.Order
}

func (m *MockWriter) SaveOrder(order model.Order) error {
	m.saved = append(m.saved, order)
	return nil
}

// MockNotifier - мок отправителя уведомлений
type MockNotifier struct {
	sentTo []string
}

func (m *MockNotifier) Send(customer string) error {
	m.sentTo = append(m.sentTo, customer)
	return nil
}

func TestCreateOrder(t *testing.T) {
	repo := &MockWriter{}
	notifier := &MockNotifier{}
	service := New(repo, notifier)

	err := service.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		t.Fatalf("CreateOrder вернул ошибку: %v", err)
	}

	if len(repo.saved) != 1 {
		t.Fatalf("ожидался 1 сохранённый заказ, получено %d", len(repo.saved))
	}
	if repo.saved[0].Customer != "Иван" {
		t.Errorf("неверный клиент в заказе: %s", repo.saved[0].Customer)
	}
	if repo.saved[0].Status != "pending" {
		t.Errorf("неверный статус заказа: %s", repo.saved[0].Status)
	}

	if len(notifier.sentTo) != 1 || notifier.sentTo[0] != "Иван" {
		t.Errorf("уведомление не отправлено клиенту: %v", notifier.sentTo)
	}
}
