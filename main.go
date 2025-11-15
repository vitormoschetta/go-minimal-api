package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check request")
	fmt.Fprintf(w, "OK")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome request")

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	log.Printf("Welcome, %s!", name)
	fmt.Fprintf(w, "Welcome, %s!", name)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", health)
	r.Get("/health", health)
	r.Get("/welcome", welcome)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
