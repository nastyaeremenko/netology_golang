package sqlite

import (
	"database/sql"

	"example/solid/internal/repository/model"
)

// SQLiteRepo - реализация репозитория для sqlite
type SQLiteRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

// Init создаёт таблицу заказов, если её ещё нет
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

// SaveOrder сохраняет заказ в базу данных
func (r *SQLiteRepo) SaveOrder(order model.Order) error {
	_, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		order.Customer, order.Products, order.Total, order.Status,
	)
	return err
}
