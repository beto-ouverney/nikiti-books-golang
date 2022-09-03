package handler

import (
	"github.com/beto-ouverney/nikiti-books/controller"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
)

// Delete is a handler that send data to the controller and returns the response to the client
func Delete(w http.ResponseWriter, r *http.Request) {
	var status int
	var response []byte
	defer r.Body.Close()

	title, errP := url.QueryUnescape(chi.URLParam(r, "title"))
	if errP != nil {
		errorReturn(w, 500, errP.Error())
	}
	c := controller.New()

	err := c.Delete(title)

	if err != nil {

		errorHandler(err, status, response, w)

	} else {
		status = 204
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		return
	}

}
