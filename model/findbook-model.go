package model

import (
	"context"
	"fmt"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindBook finds a book by your title
func FindBook(param string) (*entity.Book, *customerror.CustomError) {
	//FIND ONE
	filter := bson.D{{"title", param}}
	bookResult := entity.Book{}
	opts := options.FindOne().SetProjection(bson.D{{"created", 0}, {"updated", 0}, {"_id", 0}})

	err := collection.FindOne(context.Background(), filter, opts).Decode(&bookResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found")
			fmt.Printf("Error: %v \n", err)
			return nil, &customerror.CustomError{Code: customerror.ENOTFOUND, Op: "booksmodel.FindBook", Err: nil}
		}
		return nil, &customerror.CustomError{Code: customerror.EINVALID, Op: "booksmodel.FindBook", Err: err}
	}
	return &bookResult, nil
}
