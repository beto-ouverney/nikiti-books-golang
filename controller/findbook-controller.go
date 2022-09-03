package controller

import (
	"encoding/json"
	"github.com/beto-ouverney/nikiti-books/customerror"
)

// FindBook returns a book from service and return a json to handler
func (c *BookController) FindBook(name string) ([]byte, *customerror.CustomError) {
	book, err := c.Service.FindBook(name)
	if err != nil {
		return nil, err
	}
	bookJson, errJ := json.MarshalIndent(book, "", "    ")
	if errJ != nil {
		return nil, &customerror.CustomError{Op: "bookcontroller.FindBook", Err: errJ, Code: customerror.EINTERNAL}
	}

	return bookJson, nil
}
