package service

import (
	"errors"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
)

func validationsFields(book *entity.Book) *customerror.CustomError {
	if book == nil {
		return &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("book cannot be null")}
	}
	if book.Title == "" {
		return &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("title field is invalid")}
	}
	if book.Author == "" || len(book.Author) < 3 {
		return &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("author field is invalid, must be more than 3 characters")}
	}
	if book.Synopsis == "" || len(book.Synopsis) < 80 {
		return &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("synopsis field is invalid, must be more than 80 characters")}
	}
	if len(book.Category) < 1 {
		return &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("category field is invalid, must have at least one category")}
	}
	return nil
}
