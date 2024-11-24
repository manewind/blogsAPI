package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"os"
)

func ConnectDb() (*pgx.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file %v", err)
	}

	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil

}
