package controller

import (
	"encoding/json"
	"github.com/beto-ouverney/nikiti-books/customerror"
)

// FindAll returns all books from service and return a json
func (c *BookController) FindAll() ([]byte, *customerror.CustomError) {
	books, err := c.Service.FindAll()
	if err != nil {
		return nil, err
	}

	booksJ, errJ := json.MarshalIndent(books, "", "    ")
	if errJ != nil {
		return nil, &customerror.CustomError{Op: "bookcontroller.FindAll", Err: errJ, Code: customerror.EINTERNAL}
	}

	return booksJ, nil
}
