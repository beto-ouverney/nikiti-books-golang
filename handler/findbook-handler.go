package handler

import (
	"github.com/beto-ouverney/nikiti-books/controller"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
)

// FindBook is a function to find a book, it receives a request and response. Send request to controller and return response
func FindBook(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	defer r.Body.Close()

	title, errP := url.QueryUnescape(chi.URLParam(r, "title"))
	if errP != nil {
		errorReturn(w, 500, errP.Error())
	}
	c := controller.New()

	response, err := c.FindBook(title)

	if err != nil {

		if err.Code == customerror.ENOTFOUND {

			status = 404
			response = []byte("{\"message\":\"Book not found\"}")

		} else if err.Code == customerror.ECONFLICT {

			status = 401
			response = []byte("{\"message\":\"" + err.Error() + "\"}")

		} else {
			status = 400
			response = []byte("{\"message\":\"" + err.Error() + "\"}")
		}
	} else {
		status = 200
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, 500, errW.Error())
	}
}
