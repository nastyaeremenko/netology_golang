package repository

import "example/solid/internal/repository/model"

// RepositoryWriter - интерфейс для записи заказа в хранилище
type RepositoryWriter interface {
	SaveOrder(order model.Order) error
}

// RepositoryInitializer - отдельный интерфейс для инициализации БД (создание таблиц),
// чтобы сервис не зависел от методов, которые он не использует
type RepositoryInitializer interface {
	Init() error
}
