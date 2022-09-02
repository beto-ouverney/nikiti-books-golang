package controller

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
)

// Add is a function to add a book, it receives a title, author, synopis and category. Create entity.Book and send to service
func (m *BookController) Add(title, author, synopis string, category []string) *customerror.CustomError {
	book := &entity.Book{
		Title:    title,
		Author:   author,
		Synopsis: synopis,
		Category: category,
	}
	err := m.Service.Add(book)
	if err != nil {
		return err
	}

	return nil
}
