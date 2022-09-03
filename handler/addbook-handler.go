package handler

import (
	"encoding/json"
	"github.com/beto-ouverney/nikiti-books/controller"
	"net/http"
)

// Add is a handler that send data to the controller and returns the response to the client
func Add(w http.ResponseWriter, r *http.Request) {
	var status int
	var response []byte
	defer r.Body.Close()

	data := struct {
		Title    string   `json:"title"`
		Author   string   `json:"author"`
		Category []string `json:"category"`
		Synopsis string   `json:"synopsis"`
	}{}

	errJ := json.NewDecoder(r.Body).Decode(&data)
	if errJ != nil {
		errorReturn(w, 500, errJ.Error())
	}

	c := controller.New()

	err := c.Add(data.Title, data.Author, data.Synopsis, data.Category)
	// Tratamento de erro
	if err != nil {

		errorHandler(err, status, response, w)
	} else {
		status = 201
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		return
	}

}
