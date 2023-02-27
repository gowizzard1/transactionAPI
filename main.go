package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	config2 "transationAPI/config"
	gorm2 "transationAPI/database/gorm"
	"transationAPI/transaction"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	host, err := config2.Get("HOST")
	if err != nil {
		return
	}
	user, err := config2.Get("DB_USER")
	if err != nil {
		return
	}
	dbName, err := config2.Get("DB_NAME")
	if err != nil {
		return
	}
	dbPass, err := config2.Get("DB_PASSWORD")
	if err != nil {
		return
	}
	dbPort, err := config2.Get("DB_PORT")
	if err != nil {
		return
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, dbPass, dbName, dbPort)
	dialector := postgres.Open(dsn)
	config := gorm.Config{}
	maxOpenConns := 10
	maxIdleConns := 5
	connMaxLifetime := 300 // seconds
	db, err := gorm2.New(dialector, config, maxOpenConns, maxIdleConns, connMaxLifetime)
	if err != nil {
		panic(err)
	}
	// new repository
	repo := transaction.NewRepository(db)
	// new service
	service := transaction.NewService(*repo)
	// new controller
	controller := transaction.NewController(service)

	http.HandleFunc("/create", controller.Create)
	http.HandleFunc("/balance", controller.Balance)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
