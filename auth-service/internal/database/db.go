package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    connStr := "postgres://postgres:yourpassword@localhost:5432/authdb?sslmode=disable"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Cannot open PostgreSQL:", err)
    }
    err = db.Ping()
    if err != nil {
        log.Fatal("Cannot connect to PostgreSQL:", err)
    }

    log.Println("Connected to PostgreSQL successfully")
    DB = db
}
