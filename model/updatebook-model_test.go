package model_test

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/beto-ouverney/nikiti-books/model/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookModel_Update(t *testing.T) {
	assert := assert.New(t)

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
		msg1  string
	}{
		{
			name: "Should be able to update a book by title",
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
			msg1:  "Error should be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := new(mocks.IBookModel)
			m.On("Update", tt.args.param, tt.args.book).Return(tt.want, tt.want1)

			got, got1 := m.Update(tt.args.param, tt.args.book)
			assert.Equalf(tt.want, got, tt.msg)
			assert.Equalf(tt.want1, got1, tt.msg1)
		})
	}
}
