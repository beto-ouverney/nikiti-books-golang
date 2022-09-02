package model

import (
	"context"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindAll finds all books
func (m *BookModel) FindAll() (*[]entity.Book, *customerror.CustomError) {

	var books []entity.Book
	// Removed `created` and `updated` fields from result as they were not relevant for the final users. Nevertheless,
	// the admin user can remove them manually
	opts := options.Find().SetProjection(bson.D{{"_id", 0}})

	cur, err := collection.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, &customerror.CustomError{Code: customerror.EINTERNAL, Op: "booksmodel.FindAll", Err: err}
	}
	defer cur.Close(context.Background())

	err = cur.All(context.Background(), &books)
	if err != nil {
		return nil, &customerror.CustomError{Code: customerror.EINTERNAL, Op: "booksmodel.FindAll", Err: err}
	}

	return &books, nil
}
