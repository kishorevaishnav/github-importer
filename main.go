package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

// initialise to load environment variable from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := chi.NewRouter()

	r.Get("/", indexHandler)
	r.Post("/upload", uploadHandler)

	fmt.Println("Running Server: http://localhost:8080 ")
	http.ListenAndServe(":8080", r)
}
