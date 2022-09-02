package service_test

import (
	"errors"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	mocks_service "github.com/beto-ouverney/nikiti-books/service/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookService_FindBook(t *testing.T) {

	assertions := assert.New(t)

	type args struct {
		title string
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
				title: "The Lord of the Rings",
			},
			want: &entity.Book{
				Title:    "The Lord of the Rings",
				Author:   "J.R.R. Tolkien",
				Category: []string{"High-Fantasy", "Adventure"},
				Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
			},
			want1: nil,
			msg:   "Book should be found",
			msg1:  "Error should be nil",
		},
		{
			name: "Should not be able to find a book by title if title is empty",
			args: args{
				title: "",
			},
			want:  nil,
			want1: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.FindBook", Err: errors.New("title is invalid")},
			msg:   "Book should not be found if title is empty",
			msg1:  "Error should not be nil if title is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := mocks_service.NewIBookService(t)
			m.On("FindBook", tt.args.title).Return(tt.want, tt.want1)

			got, got1 := m.FindBook(tt.args.title)

			assertions.Equalf(tt.want, got, tt.msg)
			assertions.Equalf(tt.want1, got1, tt.msg1)
		})
	}
}
