package service

import (
	"errors"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
)

// FindBook returns a book from model by title if title is valid
func (s *BookService) FindBook(title string) (*entity.Book, *customerror.CustomError) {
	if title == "" {
		return nil, &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.FindBook", Err: errors.New("title is invalid")}
	}

	book, err := s.Model.FindBook(title)
	if err == nil {
		return book, nil
	}

	return nil, err
}
