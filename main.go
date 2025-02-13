package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check")
	fmt.Fprintf(w, "OK")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Welcome, %s!", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", health)
	r.HandleFunc("/welcome", welcome)
	http.ListenAndServe(":8080", r)
}
