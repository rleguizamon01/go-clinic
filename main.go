package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

type Medicine struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	connectToDatabase()
	r := newRouter()
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase() {
	cfg := mysql.Config{
		User:                 getDotEnvValue("DBUSER"),
		Passwd:               getDotEnvValue("DBPASS"),
		Net:                  getDotEnvValue("DBNET"),
		Addr:                 getDotEnvValue("DBADDRESS"),
		DBName:               getDotEnvValue("DBNAME"),
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/medicines", getMedicinesHandler).Methods("GET")
	r.HandleFunc("/medicines", createMedicineHandler).Methods("POST")
	return r
}

func getDotEnvValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}
