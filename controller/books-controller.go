package controller

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/service"
)

type IBookController interface {
	FindAll() ([]byte, *customerror.CustomError)
	FindBook(name string) ([]byte, *customerror.CustomError)
	Delete(name string) *customerror.CustomError
	Add(title, author, synopsis string, category []string) *customerror.CustomError
	Update(param, title, author, synopsis string, category []string) *customerror.CustomError
}

type BookController struct {
	Service service.IBookService
}

func New() *BookController {
	return &BookController{
		Service: service.New(),
	}
}
