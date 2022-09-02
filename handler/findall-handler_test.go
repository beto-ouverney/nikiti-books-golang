package handler_test

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/nikiti-books/config"
	"github.com/beto-ouverney/nikiti-books/customrouter"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/handler"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var collection *mongo.Collection

var books = []entity.Book{
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
	for _, t := range books {
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

func TestFindAll(t *testing.T) {

	assertions := assert.New(t)

	t.Log(config.MONGO_CONNECT)

	// initialize the database if in the test environment, verify the port number
	if strings.Contains(config.MONGO_CONNECT, "6306") {
		t.Log("Initializing the database for testing")
		initDBTEST(t)
		//defer dropDb(t)
	} else {
		t.Skip("Skipping test because it is not a test environment, the port number is not 6306")
	}

	router := &customrouter.Router{}

	router.Route(http.MethodGet, "/books", handler.FindAll)

	tests := []struct {
		describe        string
		books           []entity.Book
		expectedStatus  int
		expectedMessage interface{}
		assertMessage   string
	}{
		{
			describe:        "Should be able to find all books",
			books:           books,
			expectedStatus:  http.StatusOK,
			expectedMessage: books,
			assertMessage:   "Should be able to find all books",
		},
	}

	for _, test := range tests {
		t.Run(test.describe, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/books", nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assertions.Equal(test.expectedStatus, rr.Code, "Status code should be equal")
			var actual []entity.Book

			err := json.Unmarshal(rr.Body.Bytes(), &actual)
			if err != nil {
				t.Fatal(err)
			}
			assertions.Equal(test.expectedMessage, actual, test.assertMessage)
		})

	}

}
