package service

import (
	"errors"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
)

// Update updates a book from model by title if title is valid, if Book exist and have valid fields
func (s *BookService) Update(param string, book *entity.Book) *customerror.CustomError {
	if param == "" {
		return &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Update", Err: errors.New("param is invalid")}
	}

	//Verify fields
	err := validationsFields(book, "service.Update")
	if err != nil {
		return err
	}

	//Verify if book exist
	bookExist, err := s.Model.FindBook(param)
	if err != nil {
		return err
	}

	if bookExist == nil {
		return &customerror.CustomError{Code: customerror.ENOTFOUND, Op: "service.Update", Err: errors.New("book not found")}
	}

	err = s.Model.Update(param, book)
	if err != nil {
		return err
	}

	return nil
}
