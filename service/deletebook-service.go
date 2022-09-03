package service

import (
	"errors"
	"github.com/beto-ouverney/nikiti-books/customerror"
)

// Delete deletes a book from model by title if title is valid and Book exist
func (s *BookService) Delete(title string) *customerror.CustomError {
	if title == "" {
		return &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Delete", Err: errors.New("title is invalid")}
	}

	//Verify if book exist
	bookExist, err := s.Model.FindBook(title)
	if err != nil {
		return err
	}

	if bookExist == nil {
		return &customerror.CustomError{Code: customerror.ENOTFOUND, Op: "service.Delete", Err: errors.New("book not found")}
	}

	err = s.Model.Delete(title)
	if err == nil {
		return nil
	}

	return err
}
