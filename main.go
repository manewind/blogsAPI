package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

var dbConn *pgx.Conn

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// Example database query
	var message string
	var err = dbConn.QueryRow(context.Background(), `SELECT 'Hello, World from the database!'`).Scan(&message)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, message)
}

func main() {
	conn, err := ConnectDb()
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer conn.Close(context.Background())

	dbConn = conn

	http.HandleFunc("/", exampleHandler)

	// Start the HTTP server and listen on port 8080
	log.Println("Server is starting on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start", err)
	}
}
