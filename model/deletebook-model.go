package model

import (
	"context"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"go.mongodb.org/mongo-driver/bson"
)

// Delete deletes a book from the database
func (m *BookModel) Delete(param string) *customerror.CustomError {

	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"title": param})
	if err != nil {
		return &customerror.CustomError{Code: customerror.EINVALID, Op: "booksmodel.DeleteBook", Err: err}
	}

	if deleteResult.DeletedCount == 0 {
		return &customerror.CustomError{Code: customerror.ENOTFOUND, Op: "booksmodel.DeleteBook", Err: nil}
	}

	return nil
}
