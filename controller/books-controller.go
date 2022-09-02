package controller

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/service"
)

type IBookController interface {
	FindAll() ([]byte, *customerror.CustomError)
}

type BookController struct {
	Service service.IBookService
}

func New() *BookController {
	return &BookController{
		Service: service.New(),
	}
}
