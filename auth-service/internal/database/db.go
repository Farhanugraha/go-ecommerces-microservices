package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
 
    user:= os.Getenv("DB_USER")
    password:= os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    connStr := fmt.Sprintf(
    "postgres://%s:%s@%s:%s/%s?sslmode=disable",
    user, password, host, port, dbname,   
    )

    db, err:= sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Cannot open PostgreSQL:", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Cannot connect to PostgreSQL:", err)
    }

    log.Println("Connected to PostgreSQL succesfully!")
    DB = db


}
