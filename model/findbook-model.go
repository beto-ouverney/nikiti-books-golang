package model

import (
	"context"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindBook finds a book by your title
func (m *BookModel) FindBook(param string) (*entity.Book, *customerror.CustomError) {
	//FIND ONE
	filter := bson.D{{"title", param}}
	bookResult := entity.Book{}

	// Removed `created` and `updated` fields from result as they were not relevant for the final users. Nevertheless,
	// the admin user can remove them manually
	opts := options.FindOne().SetProjection(bson.D{{"_id", 0}})

	err := collection.FindOne(context.Background(), filter, opts).Decode(&bookResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &customerror.CustomError{Code: customerror.ENOTFOUND, Op: "booksmodel.FindBook", Err: nil}
		}
		return nil, &customerror.CustomError{Code: customerror.EINVALID, Op: "booksmodel.FindBook", Err: err}
	}
	return &bookResult, nil
}
