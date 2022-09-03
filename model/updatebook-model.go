package model

import (
	"context"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"go.mongodb.org/mongo-driver/bson"
)

// Update updates a book in the database
func (m *BookModel) Update(param string, book *entity.Book) *customerror.CustomError {

	filter := bson.M{"title": param}
	update := bson.M{"$set": book}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return &customerror.CustomError{Code: customerror.EINVALID, Op: "booksmodel.Update", Err: err}
	}
	return nil
}
