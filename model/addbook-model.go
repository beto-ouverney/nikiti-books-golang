package model

import (
	"context"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
)

//Add adds a book to the database
func Add(book *entity.Book) *customerror.CustomError {

	_, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		return &customerror.CustomError{Code: customerror.EINVALID, Op: "booksmodel.Add", Err: err}
	}
	return nil
}
