package handler

import (
	"encoding/json"
	"github.com/beto-ouverney/nikiti-books/controller"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
)

// Update is a function to update a book, it receives a request and response. Send request to controller and return response
func Update(w http.ResponseWriter, r *http.Request) {
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

	param, errP := url.QueryUnescape(chi.URLParam(r, "title"))
	if errP != nil {
		errorReturn(w, 500, errP.Error())
	}

	c := controller.New()

	err := c.Update(param, data.Title, data.Author, data.Synopsis, data.Category)

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
