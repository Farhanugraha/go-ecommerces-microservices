package main

import (
	"auth-service/internal/database"
	"fmt"
	"log"
	"net/http"
)
func main() {

    database.Connect()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello this auth service")
    })

    log.Println("Auth Service is running on port 8080...")
    log.Println("Hello World from Auth Service!")

    log.Fatal(http.ListenAndServe(":8080", nil))
}
