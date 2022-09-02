package model_test

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/model/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAll(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name  string
		want  *[]entity.Book
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name: "Should be able to find all books",
			want: &[]entity.Book{
				{
					Title:    "The Lord of the Rings",
					Author:   "J.R.R. Tolkien",
					Category: []string{"Fantasy", "Adventure"},
					Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
				},
			},
			want1: nil,
			msg:   "Books should be found",
			msg1:  "Error should be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := new(mocks.IBookModel)
			m.On("FindAll").Return(tt.want, nil)

			got, got1 := m.FindAll()
			assert.Equalf(tt.want, got, tt.msg)
			assert.Equalf(tt.want1, got1, tt.msg1)
		})
	}
}
