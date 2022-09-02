package model

import (
	"context"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Update updates a book in the database
func (m *BookModel) Update(param string, book *entity.Book) (*entity.Book, *customerror.CustomError) {

	book.Updated = time.Now()

	filter := bson.M{"title": param}
	update := bson.M{"$set": book}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, &customerror.CustomError{Code: customerror.EINVALID, Op: "booksmodel.Update", Err: err}
	}
	return book, nil
}
