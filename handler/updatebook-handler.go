package handler

import (
	"encoding/json"
	"fmt"
	"github.com/beto-ouverney/nikiti-books/controller"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"net/url"
)

// Update is a function to update a book, it receives a request and response. Send request to controller and return response
func Update(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	defer r.Body.Close()

	data := struct {
		Title    string   `json:"title"`
		Author   string   `json:"author"`
		Category []string `json:"category"`
		Synopsis string   `json:"synopsis"`
	}{}
	fmt.Printf("data: %v", data)
	fmt.Printf("data: %v", data)
	fmt.Printf("data: %v", data)
	errJ := json.NewDecoder(r.Body).Decode(&data)
	if errJ != nil {
		errorReturn(w, r, 500, errJ.Error())
	}

	param, errP := url.QueryUnescape(chi.URLParam(r, "title"))
	if errP != nil {
		errorReturn(w, r, 500, errP.Error())
	}

	c := controller.New()

	err := c.Update(param, data.Title, data.Author, data.Synopsis, data.Category)

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
		status = 200
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
