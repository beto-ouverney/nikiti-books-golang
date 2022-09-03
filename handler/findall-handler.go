package handler

import (
	"github.com/beto-ouverney/nikiti-books/controller"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"log"
	"net/http"
)

// FindAll is a function to find all books, it receives a request and response. Send request to controller and return response
func FindAll(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	defer r.Body.Close()

	c := controller.New()

	response, err := c.FindAll()

	if err != nil {
		status = 400
		response = []byte("{\"message\":\"" + err.Message + "\"}")
		if err.Code != customerror.ECONFLICT {
			log.Printf("Error: %v", err)
			response = []byte("{\"message\":\"" + err.Message + "\"}")
		}
	} else {
		status = 200
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, r, 500, errW.Error())
	}
}
