package handler

import (
	"github.com/beto-ouverney/nikiti-books/controller"
	"net/http"
)

// FindAll is a function to find all books, it receives a request and response. Send request to controller and return response
func FindAll(w http.ResponseWriter, r *http.Request) {
	var status int
	var response []byte
	defer r.Body.Close()

	c := controller.New()

	response, err := c.FindAll()

	if err != nil {
		errorHandler(err, status, response, w)
	}

	status = 200
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, 500, errW.Error())
	}
}
