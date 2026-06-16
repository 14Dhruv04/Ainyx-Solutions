package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connStr := os.Getenv("DB_URL")

	if connStr == "" {
		connStr = "postgres://postgres:YOUR_PASSWORD@localhost:5432/userdb?sslmode=disable"
	}

	return sql.Open("postgres", connStr)
}