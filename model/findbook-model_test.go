package model_test

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/model/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindBook(t *testing.T) {

	assert := assert.New(t)

	type args struct {
		param string
	}
	tests := []struct {
		name  string
		args  args
		want  *entity.Book
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name: "Should be able to find a book by title",
			args: args{
				param: "The Lord of the Rings",
			},
			want: &entity.Book{
				Title:    "The Lord of the Rings",
				Author:   "J.R.R. Tolkien",
				Category: []string{"Fantasy", "Adventure"},
				Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
			},
			want1: nil,
			msg:   "Book should be found",
			msg1:  "Error should be nil",
		},
		{
			name: "Should not be able to find a book by title",
			args: args{
				param: "The Lord of Code",
			},
			want:  nil,
			want1: &customerror.CustomError{Code: customerror.ENOTFOUND, Op: "booksmodel.FindBook", Err: nil},
			msg:   "Book should not be found",
			msg1:  "Error should not be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := new(mocks.IBookModel)
			m.On("FindBook", tt.args.param).Return(tt.want, tt.want1)

			got, got1 := m.FindBook(tt.args.param)
			assert.Equalf(tt.want, got, tt.msg)
			assert.Equalf(tt.want1, got1, tt.msg1)
		})
	}
}
