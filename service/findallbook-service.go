package service

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
)

// FindAll returns all books from model
func (s *BookService) FindAll() (*[]entity.Book, *customerror.CustomError) {
	books, err := s.Model.FindAll()
	if err == nil {
		return books, nil
	}
	return nil, err
}
