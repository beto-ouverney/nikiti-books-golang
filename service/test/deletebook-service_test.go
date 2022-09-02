package test_test

import (
	"errors"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	mocks_service "github.com/beto-ouverney/nikiti-books/service/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookService_Delete(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		title string
	}
	tests := []struct {
		name  string
		args  args
		want  *entity.Book
		want2 *customerror.CustomError
		want1 *customerror.CustomError
		msg   string
	}{
		{
			name: "Should be able to delete a book by title",
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
			msg:   "Book should be deleted",
		},
		{
			name: "Should not be able to delete a book by title if title is empty",
			args: args{
				title: "",
			},
			want:  nil,
			want1: &customerror.CustomError{Code: customerror.ECONFLICT, Op: "service.Delete", Err: errors.New("title is invalid")},
			msg:   "Book should not be deleted if title is empty",
		},
		{
			name: "Should not be able to delete a book if book does not exist",
			args: args{
				title: "The Lord ",
			},
			want:  nil,
			want1: &customerror.CustomError{Code: customerror.ENOTFOUND, Op: "service.Delete", Err: errors.New("book not found")},
			msg:   "Book should not be deleted if book does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks_service.NewIBookService(t)
			m.On("Delete", tt.args.title).Return(tt.want1)

			got := m.Delete(tt.args.title)
			assertions.Equal(tt.want1, got, tt.msg)
		})
	}
}
