package service_test

import (
	"errors"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	mocks_service "github.com/beto-ouverney/nikiti-books/service/mocks"
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
			m := mocks_service.NewIBookService(t)
			m.On("Update", tt.args.param, tt.args.book).Return(tt.want1)

			got := m.Update(tt.args.param, tt.args.book)
			assertions.Equalf(tt.want1, got, tt.msg)
		})
	}
}

func TestBookService_UpdateValidationsFields(t *testing.T) {

	assertions := assert.New(t)

	type args struct {
		param string
		book  *entity.Book
	}
	tests := []struct {
		name string
		args args
		want *customerror.CustomError
		msg  string
	}{
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
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Update", Err: errors.New("title field is invalid")},
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
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Update", Err: errors.New("author field is invalid, must be more than 3 characters")},
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
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Update", Err: errors.New("synopsis field is invalid, must be more than 80 characters")},
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
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Update", Err: errors.New("category field is invalid, must have at least one category")},
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
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Update", Err: errors.New("author field is invalid, must be more than 3 characters")},
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
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Update", Err: errors.New("synopsis field is invalid, must be more than 80 characters")},
			msg:  "Book should not be added if synopsis field don`t have more than 80 characters",
		},
		{
			name: "Shoud not be able to add a book if book is nil",
			args: args{
				book: nil,
			},
			want: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Update", Err: errors.New("book cannot be null")},
			msg:  "Book should not be added if book is nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks_service.NewIBookService(t)
			m.On("Update", tt.args.param, tt.args.book).Return(tt.want)

			got := m.Update(tt.args.param, tt.args.book)
			assertions.Equalf(tt.want, got, tt.msg)
		})
	}
}
