package test_test

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/model/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		book *entity.Book
	}

	tests := []struct {
		name string
		args args
		want *customerror.CustomError
	}{
		{
			name: "Should be able to Add a book",
			args: args{
				book: &entity.Book{
					Title:    "The Lord of the Rings",
					Author:   "J.R.R. Tolkien",
					Category: []string{"Fantasy", "Adventure"},
					Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
					Created:  time.Now(),
					Updated:  time.Now(),
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := new(mocks.IBookModel)

			m.On("Add", tt.args.book).Return(nil)

			err := m.Add(tt.args.book)

			assert.Nil(err, "Error should be nil")
		})
	}
}
