package controller

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
)

// Update is a function to update a book, it receives a param, title, author, synopsis and category.
//Create entity.Book and send to service with param to find the book
func (m *BookController) Update(param, title, author, synopsis string, category []string) *customerror.CustomError {
	book := &entity.Book{
		Title:    title,
		Author:   author,
		Synopsis: synopsis,
		Category: category,
	}

	err := m.Service.Update(param, book)
	if err != nil {
		return err
	}
	return nil
}
