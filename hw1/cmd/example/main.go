package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"example/solid/internal/notifier"
	"example/solid/internal/repository/sqlite"
	"example/solid/internal/service"
)

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := sqlite.New(db)
	err = repo.Init()
	if err != nil {
		log.Fatal(err)
	}

	// сервис с отправкой по email
	emailService := service.New(repo, &notifier.EmailSender{})
	err = emailService.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}

	// тот же сервис, но с отправкой по sms - код сервиса не меняется
	smsService := service.New(repo, &notifier.SMSSender{})
	err = smsService.CreateOrder("Мария", []string{"orange", "grape"}, 7.2)
	if err != nil {
		log.Fatal(err)
	}
}
