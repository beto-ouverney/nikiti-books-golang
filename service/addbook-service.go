package service

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
)

// Add adds a book to model if book is not nil, and have valid fields
func (s *BookService) Add(book *entity.Book) *customerror.CustomError {
	//Verify fields before add
	err := validationsFields(book)
	if err != nil {
		return err
	}

	/* I don't need to verify if book exist because I suppose he's allowed to own the same books
	bookExist, err := s.Model.FindBook(book.Title)
	if err != nil {
		return err
	}
	*/

	err = s.Model.Add(book)
	if err == nil {
		return nil
	}
	return err
}
