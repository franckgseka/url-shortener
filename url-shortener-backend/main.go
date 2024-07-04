package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/teris-io/shortid"
)

type URL struct {
	ID       string `json:"id"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
	Clicks   int    `json:"clicks"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTable()

	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortenURL).Methods("POST")
	r.HandleFunc("/{id}", redirect).Methods("GET")
	r.HandleFunc("/stats", stats).Methods("GET")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createTable() {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS urls (id TEXT PRIMARY KEY, long_url TEXT, short_url TEXT, clicks INTEGER)")
	if err != nil {
		log.Fatalf("Failed to prepare table creation statement: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("Failed to execute table creation statement: %v", err)
	}
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	var url URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := shortid.Generate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shortURL := "http://localhost:8080/" + id

	url.ID = id
	url.ShortURL = shortURL
	url.Clicks = 0

	statement, err := db.Prepare("INSERT INTO urls (id, long_url, short_url, clicks) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = statement.Exec(url.ID, url.LongURL, url.ShortURL, url.Clicks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(url)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var url URL
	row := db.QueryRow("SELECT id, long_url, short_url, clicks FROM urls WHERE id = ?", id)
	err := row.Scan(&url.ID, &url.LongURL, &url.ShortURL, &url.Clicks)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	url.Clicks++
	statement, err := db.Prepare("UPDATE urls SET clicks = ? WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = statement.Exec(url.Clicks, url.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url.LongURL, http.StatusMovedPermanently)
}

func stats(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, long_url, short_url, clicks FROM urls")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var urls []URL
	for rows.Next() {
		var url URL
		err := rows.Scan(&url.ID, &url.LongURL, &url.ShortURL, &url.Clicks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		urls = append(urls, url)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(urls)
}
