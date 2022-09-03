package handler_test

import (
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

func TestDelete(t *testing.T) {

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

	router := chi.NewRouter()

	router.Delete("/books/{title}", handler.Delete)

	tests := []struct {
		describe        string
		books           entity.Book
		title           string
		expectedStatus  int
		expectedMessage interface{}
		assertMessage   string
	}{
		{
			describe:        "Should be able to delete a book",
			title:           booksMock[0].Title,
			expectedStatus:  http.StatusNoContent,
			expectedMessage: []byte(nil),
			assertMessage:   "Should be able to delete a book",
		},
		{
			describe:        "Should not be able to delete a book",
			title:           booksMock[0].Title,
			expectedStatus:  http.StatusNotFound,
			expectedMessage: []byte("{\"message\":\"Book not found\"}"),
			assertMessage:   "Should not be able to delete a book",
		},
	}

	for _, test := range tests {
		t.Run(test.describe, func(t *testing.T) {
			path := fmt.Sprintf("/books/%s", url.QueryEscape(test.title))
			t.Log(url.QueryEscape(test.title))
			req := httptest.NewRequest(http.MethodDelete, path, nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assertions.Equal(test.expectedStatus, rr.Code, "Status code should be equal")

			assertions.Equal(test.expectedMessage, rr.Body.Bytes(), test.assertMessage)

		})

	}
}
