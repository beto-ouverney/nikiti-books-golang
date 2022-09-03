package test_test

import (
	"encoding/json"
	"github.com/beto-ouverney/nikiti-books/config"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/handler"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFindAll(t *testing.T) {

	assertions := assert.New(t)

	t.Log(config.MONGO_CONNECT)

	// initialize the database if in the test environment, verify the port number
	if strings.Contains(config.MONGO_CONNECT, "6306") {
		t.Log("Initializing the database for testing")
		initDBTEST(t)
		defer dropDb(t)
	} else {
		t.Skip("Skipping test because it is not a test environment, the port number is not 6306")
	}

	router := chi.NewRouter()

	router.Get("/books", handler.FindAll)

	tests := []struct {
		describe        string
		books           []entity.Book
		expectedStatus  int
		expectedMessage interface{}
		assertMessage   string
	}{
		{
			describe:        "Should be able to find all books",
			books:           booksMock,
			expectedStatus:  http.StatusOK,
			expectedMessage: booksMock,
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
