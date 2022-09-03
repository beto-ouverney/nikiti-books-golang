package model

import (
	"context"
	"fmt"
	"github.com/beto-ouverney/nikiti-books/config"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection

// IBookModel is the interface for the BookModel
type IBookModel interface {
	Add(book *entity.Book) *customerror.CustomError
	FindBook(param string) (*entity.Book, *customerror.CustomError)
	FindAll() (*[]entity.Book, *customerror.CustomError)
	Delete(param string) *customerror.CustomError
	Update(param string, book *entity.Book) *customerror.CustomError
}

// BookModel is the model
type BookModel struct {
	IBookModel
}

// New creates a new BookModel
func New() IBookModel {

	var cred options.Credential

	cred.Username = config.MONGO_USER
	cred.Password = config.MONGO_PASSWORD

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MONGO_CONNECT).SetAuth(cred))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	collection = client.Database("nikiti_books_db").Collection("books")
	fmt.Println("Collection is ready")

	return &BookModel{}

}
