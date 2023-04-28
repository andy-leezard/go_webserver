package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDB("urlshortener.db")
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/create", CreateShortURLHandler(db)).Methods("POST")
	router.HandleFunc("/{shortURL}", RedirectHandler(db))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	if err := createTable(db); err != nil {
		log.Fatal(err)
	}

	return db
}
