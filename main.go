package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Order struct {
	ID       int
	Customer string
	Products string
	Total    float64
	Status   string
}

// RepositoryWriter - интерфейс для записи заказа в хранилище
type RepositoryWriter interface {
	SaveOrder(order Order) error
}

// RepositoryInitializer - отдельный интерфейс для инициализации БД (создание таблиц)
type RepositoryInitializer interface {
	Init() error
}

// Notifier - интерфейс для отправки уведомлений
type Notifier interface {
	Send(customer string) error
}

// SQLiteRepo - реализация репозитория для sqlite
type SQLiteRepo struct {
	db *sql.DB
}

func (r *SQLiteRepo) Init() error {
	_, err := r.db.Exec(`
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer TEXT NOT NULL,
		products TEXT NOT NULL,
		total REAL NOT NULL,
		status TEXT NOT NULL
	)`)
	return err
}

func (r *SQLiteRepo) SaveOrder(order Order) error {
	_, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		order.Customer, order.Products, order.Total, order.Status,
	)
	return err
}

// EmailSender отправляет уведомления по почте
type EmailSender struct{}

func (s *EmailSender) Send(customer string) error {
	fmt.Printf("Email-уведомление отправлено клиенту %s\n", customer)
	return nil
}

// SMSSender отправляет уведомления по смс
type SMSSender struct{}

func (s *SMSSender) Send(customer string) error {
	fmt.Printf("SMS-уведомление отправлено клиенту %s\n", customer)
	return nil
}

// OrderService - бизнес-логика, зависит только от интерфейсов
type OrderService struct {
	repo     RepositoryWriter
	notifier Notifier
}

func NewOrderService(repo RepositoryWriter, notifier Notifier) *OrderService {
	return &OrderService{repo: repo, notifier: notifier}
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	order := Order{
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

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := &SQLiteRepo{db: db}
	err = repo.Init()
	if err != nil {
		log.Fatal(err)
	}

	// сервис с отправкой по email
	emailService := NewOrderService(repo, &EmailSender{})
	err = emailService.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}

	// тот же сервис, но с отправкой по sms - код сервиса не меняется
	smsService := NewOrderService(repo, &SMSSender{})
	err = smsService.CreateOrder("Мария", []string{"orange", "grape"}, 7.2)
	if err != nil {
		log.Fatal(err)
	}
}