package service

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/model"
)

type IBookService interface {
	FindAll() (*[]entity.Book, *customerror.CustomError)
	FindBook(title string) (*entity.Book, *customerror.CustomError)
	Delete(param string) *customerror.CustomError
	//Add(book *entity.Book) *customerror.CustomError
	//Delete(param string) *customerror.CustomError
	//Update(param string, book *entity.Book) (*entity.Book, *customerror.CustomError)
}

type BookService struct {
	Model model.IBookModel
}

func New() *BookService {
	return &BookService{
		Model: model.New(),
	}
}
