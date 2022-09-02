package service

import (
	"errors"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/model/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookService_Update(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		param string
		book  *entity.Book
	}
	tests := []struct {
		name  string
		args  args
		want  *entity.Book
		want1 *customerror.CustomError
		msg   string
	}{
		{
			name: "Should be able to update a book",
			args: args{
				param: "The Lord of the Rings",
				book: &entity.Book{
					Title:    "The Lord of the Rings",
					Author:   "J.R.R. Tolkien",
					Category: []string{"High-Fantasy", "Adventure"},
					Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
				},
			},
			want: &entity.Book{
				Title:    "The Lord of the Rings",
				Author:   "J.R.R. Tolkien",
				Category: []string{"High-Fantasy", "Adventure"},
				Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
			},
			want1: nil,
			msg:   "Book should be updated",
		},
		{
			name: "Should not be able to update a book if book not exist",
			args: args{
				param: "The Lord of the Rings",
				book: &entity.Book{
					Title:    "The Lord of the Rings",
					Author:   "J.R.R. Tolkien",
					Category: []string{"High-Fantasy", "Adventure"},
					Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
				},
			},
			want:  nil,
			want1: &customerror.CustomError{Code: customerror.ENOTFOUND, Op: "service.Update", Err: errors.New("book not found")},
			msg:   "Book should not be updated if book exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := new(mocks.IBookModel)
			m.On("FindBook", tt.args.param).Return(tt.want, nil)
			m.On("Update", tt.args.param, tt.args.book).Return(tt.want1)

			s := BookService{Model: m}

			got := s.Update(tt.args.param, tt.args.book)
			assertions.Equalf(tt.want1, got, tt.msg)
		})
	}
}
