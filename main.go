package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// the underscore is a blank identifier.

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

func createTable(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_url TEXT NOT NULL,
		long_url TEXT NOT NULL
	);`

	_, err := db.Exec(createTableSQL)
	return err
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Render home page with a form for URL submission
}

func CreateShortURLHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle URL submission and generate short URL
	}
}

func RedirectHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Redirect to the long URL corresponding to the short URL
	}
}