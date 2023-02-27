package main

import (
	"gorm.io/gorm"
	"log"
	"net/http"
	"sync"
	gorm2 "transationAPI/database/gorm"
	"transationAPI/transaction"
)

func main() {

	dsn := "host=localhost user=postgres password=secret dbname=mydatabase port=5432 sslmode=disable TimeZone=UTC"
	dialector := postgres.Open(dsn)
	config := gorm.Config{}
	maxOpenConns := 10
	maxIdleConns := 5
	connMaxLifetime := 300 // seconds
	db, err := gorm2.New(dialector, config, maxOpenConns, maxIdleConns, connMaxLifetime)
	if err != nil {
		panic(err)
	}
	// starting balance
	bal := 1000.0

	// new repository
	repo := transaction.NewRepository(db)
	// new service
	service := transaction.NewService(bal, &sync.Mutex{})
	// new controller
	controller := transaction.NewController(service)

	http.HandleFunc("/debit", controller.Debit)
	http.HandleFunc("/credit", controller.Credit)
	http.HandleFunc("/balance", controller.Balance)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
