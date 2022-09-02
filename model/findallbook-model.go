package model

import (
	"context"
	"fmt"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// FindAll finds all books
func (m *BookModel) FindAll() (*[]entity.Book, *customerror.CustomError) {
	var books []entity.Book

	// Removed fields created and updated from the result because they are not necessary for users final but if admin wants to see them, he can
	opts := options.Find().SetProjection(bson.D{{"created", 0}, {"updated", 0}, {"_id", 0}})

	cur, err := collection.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, &customerror.CustomError{Code: customerror.EINTERNAL, Op: "booksmodel.FindAll", Err: err}
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {

		book := entity.Book{}

		err := cur.Decode(&book)
		if err != nil {
			return nil, &customerror.CustomError{Code: customerror.EINVALID, Op: "booksmodel.FindAll", Err: err}
		}

		log.Printf("Found a single document: %+v", book)

		books = append(books, book)
	}
	fmt.Println("Books: ", books)
	return &books, nil
}
