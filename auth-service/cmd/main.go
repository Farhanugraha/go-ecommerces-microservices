package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from Auth Service")
    })

    log.Println("Auth Service is running on port 8080...")
    log.Println("Hello World from Auth Service!")

    log.Fatal(http.ListenAndServe(":8080", nil))
}
