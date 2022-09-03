package handler

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"log"
	"net/http"
)

//Devida as caracteriticas deste projeto e a construção do customerror,
//crei este handlr para tratamento de erros das camamdas inferiores

func errorHandler(err *customerror.CustomError, status int, response []byte, w http.ResponseWriter) {
	if err.Code == customerror.ENOTFOUND {

		status = 404
		response = []byte("{\"message\":\"Book not found\"}")

	} else if err.Code == customerror.ECONFLICT {
		status = 400
		log.Printf("Error: %v", err)
		response = []byte("{\"message\":\"" + err.Error() + "\"}")

	} else {
		log.Println(err)
		status = 500
		response = []byte("{\"message\":\"" + err.Error() + "\"}")
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, 500, errW.Error())
	}

}
