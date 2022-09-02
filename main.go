package main

import (
	"fmt"
	"github.com/beto-ouverney/nikiti-books/config"
	"github.com/beto-ouverney/nikiti-books/customrouter"
	"github.com/beto-ouverney/nikiti-books/handler"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	router := &customrouter.Router{}
	router.Route(http.MethodGet, "/books", handler.FindAll)

	log.Println("Server running on port " + config.PORT)
	http.ListenAndServe(config.PORT, router)

}
