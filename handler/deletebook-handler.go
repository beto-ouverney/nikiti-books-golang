package handler

import (
	"github.com/beto-ouverney/nikiti-books/controller"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"net/url"
)

// Delete is a handler that send data to the controller and returns the response to the client
func Delete(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	defer r.Body.Close()

	title, errP := url.QueryUnescape(chi.URLParam(r, "title"))
	if errP != nil {
		errorReturn(w, r, 500, errP.Error())
	}
	c := controller.New()

	err := c.Delete(title)

	if err != nil {

		if err.Code == customerror.ENOTFOUND {

			status = 404
			response = []byte("{\"message\":\"Book not found\"}")

		} else if err.Code == customerror.ECONFLICT {
			status = 400
			log.Printf("Error: %v", err)
			response = []byte("{\"message\":\"" + err.Error() + "\"}")

		} else {

			response = []byte("{\"message\":\"" + err.Error() + "\"}")
		}
	} else {
		status = 204
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
