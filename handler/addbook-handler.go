package handler

import (
	"encoding/json"
	"github.com/beto-ouverney/nikiti-books/controller"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"log"
	"net/http"
)

// Add is a handler that send data to the controller and returns the response to the client
func Add(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	defer r.Body.Close()

	data := struct {
		Title    string   `json:"title"`
		Author   string   `json:"author"`
		Category []string `json:"category"`
		Synopsis string   `json:"synopsis"`
	}{}

	errJ := json.NewDecoder(r.Body).Decode(&data)
	if errJ != nil {
		errorReturn(w, r, 500, errJ.Error())
	}

	c := controller.New()

	err := c.Add(data.Title, data.Author, data.Synopsis, data.Category)

	if err != nil {

		if err.Code == customerror.ENOTFOUND {

			status = 404
			response = []byte("{\"message\":\"Book not found\"}")

		} else if err.Code == customerror.ECONFLICT {
			status = 400
			log.Printf("Error: %v", err)
			response = []byte("{\"message\":\"" + err.Error() + "\"}")

		} else {
			status = 400
			response = []byte("{\"message\":\"" + err.Error() + "\"}")
		}
	} else {
		status = 201
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, r, 500, errW.Error())
	}
}
