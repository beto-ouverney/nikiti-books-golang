package main

import (
	"fmt"
	"github.com/beto-ouverney/nikiti-books/config"
	"github.com/beto-ouverney/nikiti-books/handler"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	r := chi.NewRouter()
	r.Get("/books", handler.FindAll)
	r.Get("/books/{title}", handler.FindBook)
	r.Delete("/books/{title}", handler.Delete)
	r.Post("/books", handler.Add)

	log.Println("Server running on port " + config.PORT)
	http.ListenAndServe(config.PORT, r)

}
