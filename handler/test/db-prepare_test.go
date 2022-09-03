package test_test

import (
	"context"
	"github.com/beto-ouverney/nikiti-books/config"
	"github.com/beto-ouverney/nikiti-books/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

var collection *mongo.Collection

var booksMock = []entity.Book{
	{
		Title:    "The Lord of the Rings",
		Author:   "J. R. R. Tolkien",
		Category: []string{"Fantasy", "Adventure"},
		Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
	},
	{
		Title:    "The Hobbit",
		Author:   "J. R. R. Tolkien",
		Category: []string{"Fantasy", "Adventure"},
		Synopsis: "The Hobbit is a children's fantasy novel by English author J. R. R. Tolkien.",
	},
}

// initDBTEST is a function that initializes the database for testing
func initDBTEST(t *testing.T) {
	var cred options.Credential

	cred.Username = config.MONGO_USER
	cred.Password = config.MONGO_PASSWORD

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MONGO_CONNECT).SetAuth(cred))
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("nikiti_books_db").Collection("books")

	// drop the database before inserting the test data
	dropDb(t)

	var ui []interface{}
	for _, t := range booksMock {
		ui = append(ui, t)
	}
	t.Log("Inserting test data")
	_, err = collection.InsertMany(context.Background(), ui)
	if err != nil {
		log.Println(err)
	}

}

// dropDb is a function that drops the database after testing
func dropDb(t *testing.T) {
	t.Log("Dropping the database...")
	err := collection.Drop(context.Background())
	if err != nil {
		log.Println(err)
	}
}
