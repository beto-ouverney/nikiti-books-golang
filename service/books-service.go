package service

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/model"
)

// IBookService is a struct that implements IBookService interface
type IBookService interface {
	FindAll() (*[]entity.Book, *customerror.CustomError)
	FindBook(title string) (*entity.Book, *customerror.CustomError)
	Delete(param string) *customerror.CustomError
	Add(book *entity.Book) *customerror.CustomError
	Update(param string, book *entity.Book) *customerror.CustomError
}

// BookService struct
type BookService struct {
	Model model.IBookModel
}

// New returns a new BookService instance
func New() *BookService {
	return &BookService{
		Model: model.New(),
	}
}
