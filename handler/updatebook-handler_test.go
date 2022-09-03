package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/beto-ouverney/nikiti-books/config"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/handler"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUpdate(t *testing.T) {
	assertions := assert.New(t)

	// initialize the database if in the test environment, verify the port number
	if strings.Contains(config.MONGO_CONNECT, "6306") {
		t.Log("Initializing the database for testing")
		initDBTEST(t)
		//defer dropDb(t)

	} else {
		t.Skip("Skipping test because it is not a test environment, the port number is not 6306")
	}

	router := chi.NewRouter()

	router.Put("/books/{title}", handler.Update)
	router.Get("/books/{title}", handler.FindBook)

	tests := []struct {
		describe        string
		book1           entity.Book
		book2           entity.Book
		expectedStatus  int
		expectedMessage interface{}
		assertMessage   string
	}{
		{
			describe:        "Should be able to update a book",
			book1:           booksMock[0],
			book2:           booksMock[1],
			expectedStatus:  200,
			expectedMessage: []byte(nil),
			assertMessage:   "Should be able to update a book",
		},
	}

	for _, test := range tests {
		t.Run(test.describe, func(t *testing.T) {
			data, err := json.Marshal(test.book1)
			if err != nil {
				t.Fatal(err)
			}

			path := fmt.Sprintf("/books/%s", url.QueryEscape(test.book2.Title))
			req := httptest.NewRequest(http.MethodPut, path, bytes.NewBuffer(data))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assertions.Equal(test.expectedStatus, rr.Code, "Status code should be equal")

			assertions.Equal(test.expectedMessage, rr.Body.Bytes(), test.assertMessage)

			path2 := fmt.Sprintf("/books/%s", url.QueryEscape(test.book1.Title))

			req2 := httptest.NewRequest(http.MethodGet, path2, nil)
			rr2 := httptest.NewRecorder()
			router.ServeHTTP(rr2, req2)

			assertions.Equal(200, 200, "Status code should be equal")
			var actual entity.Book

			err2 := json.Unmarshal(rr2.Body.Bytes(), &actual)
			if err2 != nil {
				t.Fatal(err2)
			}
			assertions.Equal(test.book1, actual, test.assertMessage)
		})

	}

}

func TestUpdateFailTests(t *testing.T) {
	assertions := assert.New(t)

	// initialize the database if in the test environment, verify the port number
	if strings.Contains(config.MONGO_CONNECT, "6306") {
		t.Log("Initializing the database for testing")
		initDBTEST(t)
		//defer dropDb(t)

	} else {
		t.Skip("Skipping test because it is not a test environment, the port number is not 6306")
	}

	router := chi.NewRouter()

	router.Put("/books/{title}", handler.Update)
	router.Get("/books/{title}", handler.FindBook)

	tests := []struct {
		describe        string
		book            entity.Book
		expectedStatus  int
		expectedMessage interface{}
		assertMessage   string
	}{
		{
			describe: "Should not be able to update a book with an empty title",
			book: entity.Book{
				Title:    "",
				Author:   "Alberto Paz",
				Category: []string{"Programming", "Golang"},
				Synopsis: "Ë um livro que fala sobre programação em Golang e a vontade de trabalhar na Taghos Tecnologia",
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: []byte(`{"message":"title field is invalid"}`),
			assertMessage:   "Should not be able to update a book with an empty title",
		},
		{
			describe: "Should not be able to update a book with an empty author",
			book: entity.Book{
				Title:    "Nikiti Books",
				Author:   "",
				Category: []string{"Programming", "Golang"},
				Synopsis: "Ë um livro que fala sobre programação em Golang e a vontade de trabalhar na Taghos Tecnologia",
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: []byte(`{"message":"author field is invalid, must be more than 3 characters"}`),
			assertMessage:   "Should not be able to update a book with an empty author",
		},
		{
			describe: "Should not be able to update a book with an empty synopsis",
			book: entity.Book{
				Title:    "Nikiti Books",
				Author:   "Alberto Paz",
				Category: []string{"Programming", "Golang"},
				Synopsis: "",
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: []byte(`{"message":"synopsis field is invalid, must be more than 30 characters"}`),
			assertMessage:   "Should not be able to update a book with an empty synopsis",
		},
		{
			describe: "Should not be able to update a book with an empty category",
			book: entity.Book{
				Title:    "Nikiti Books",
				Author:   "Alberto Paz",
				Category: []string{},
				Synopsis: "Ë um livro que fala sobre programação em Golang e a vontade de trabalhar na Taghos Tecnologia",
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: []byte(`{"message":"category field is invalid, must have at least one category"}`),
			assertMessage:   "Should not be able to update a book with an empty category",
		},
		{
			describe: "Should not be able to update a book with author with less than 3 characters",
			book: entity.Book{
				Title:    "Nikiti Books",
				Author:   "Al",
				Category: []string{"Programming", "Golang"},
				Synopsis: "Ë um livro que fala sobre programação em Golang e a vontade de trabalhar na Taghos Tecnologia",
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: []byte(`{"message":"author field is invalid, must be more than 3 characters"}`),
		},
		{
			describe: "Should not be able to update a book with title with a Synopsis with less than 30 characters",
			book: entity.Book{
				Title:    "Nikiti Books",
				Author:   "Alberto Paz",
				Category: []string{"Programming", "Golang"},
				Synopsis: "Taghos Tecnologia",
			},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: []byte(`{"message":"synopsis field is invalid, must be more than 30 characters"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.describe, func(t *testing.T) {
			data, err := json.Marshal(test.book)
			if err != nil {
				t.Fatal(err)
			}

			path := fmt.Sprintf("/books/%s", url.QueryEscape("Taghos Tecnologia"))
			req := httptest.NewRequest(http.MethodPut, path, bytes.NewBuffer(data))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assertions.Equal(test.expectedStatus, rr.Code, "Status code should be equal")

			assertions.Equal(test.expectedMessage, rr.Body.Bytes(), test.assertMessage)

		})

	}

}
