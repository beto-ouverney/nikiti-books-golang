package service

import (
	"errors"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/model/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookService_Add(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		book *entity.Book
	}
	tests := []struct {
		name string
		args args
		want *customerror.CustomError
		msg  string
	}{
		{
			name: "Should be able to add a book",
			args: args{
				book: &entity.Book{
					Title:    "The Lord of the Rings",
					Author:   "J.R.R. Tolkien",
					Category: []string{"High-Fantasy", "Adventure"},
					Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
				},
			},
			want: nil,
			msg:  "Book should be added",
		},
		{
			name: "Should not be able to add a book if title is empty",
			args: args{
				book: &entity.Book{
					Title:    "",
					Author:   "J.R.R. Tolkien",
					Category: []string{"High-Fantasy", "Adventure"},
					Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
				},
			},
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("title field is invalid")},
			msg:  "Book should not be added if title is empty",
		},
		{
			name: "Should not be able to add a book if author is empty",
			args: args{
				book: &entity.Book{
					Title:    "The Lord of the Rings",
					Author:   "",
					Category: []string{"High-Fantasy", "Adventure"},
					Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
				},
			},
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("author field is invalid, must be more than 3 characters")},
			msg:  "Book should not be added if author is empty",
		},
		{
			name: "Should not be able to add a book if synopsis is empty",
			args: args{
				book: &entity.Book{
					Title:    "The Lord of the Rings",
					Author:   "J.R.R. Tolkien",
					Category: []string{"High-Fantasy", "Adventure"},
					Synopsis: "",
				},
			},
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("synopsis field is invalid, must be more than 80 characters")},
			msg:  "Book should not be added if synopsis is empty",
		},
		{
			name: "Should not be able to add a book if category don`t have at least one category",
			args: args{
				book: &entity.Book{
					Title:    "The Lord of the Rings",
					Author:   "J.R.R. Tolkien",
					Category: []string{},
					Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
				},
			},
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("category field is invalid, must have at least one category")},
			msg:  "Book should not be added if category don`t have at least one category",
		},
		{
			name: "Should not be able to add a book if author field don`t have more than 2 characters",
			args: args{
				book: &entity.Book{
					Title:    "The Lord of the Rings",
					Author:   "JF",
					Category: []string{"High-Fantasy", "Adventure"},
					Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
				},
			},
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("author field is invalid, must be more than 3 characters")},
			msg:  "Book should not be added if author field don`t have more than 2 characters",
		},
		{
			name: "Should not be able to add a book if synopsis field don`t have more than 80 characters",
			args: args{
				book: &entity.Book{
					Title:    "The Lord of the Rings",
					Author:   "J.R.R. Tolkien",
					Category: []string{"High-Fantasy", "Adventure"},
					Synopsis: "Where is the synopsis?",
				},
			},
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("synopsis field is invalid, must be more than 80 characters")},
			msg:  "Book should not be added if synopsis field don`t have more than 80 characters",
		},
		{
			name: "Shoud not be able to add a book if book is nil",
			args: args{
				book: nil,
			},
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Add", Err: errors.New("book cannot be null")},
			msg:  "Book should not be added if book is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := new(mocks.IBookModel)
			m.On("Add", tt.args.book).Return(tt.want)

			s := BookService{Model: m}
			got := s.Add(tt.args.book)

			assertions.Equalf(tt.want, got, tt.msg)
		})
	}
}
