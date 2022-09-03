package test

import (
	mocks_controller "github.com/beto-ouverney/nikiti-books/controller/mocks"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookController_Update(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		param    string
		title    string
		author   string
		synopsis string
		category []string
	}
	tests := []struct {
		name string
		args args
		want *customerror.CustomError
		msg  string
	}{
		{
			name: "Should be able to update a book by title",
			args: args{
				param:  "The Lord of the Rings",
				title:  "The Lord of the Rings",
				author: "J. R. R. Tolkien",
				synopsis: "The Lord of the Rings is an epic high fantasy novel written by English author and scholar J. R. R. Tolkien. " +
					"The story began as a sequel to Tolkien's 1937 fantasy novel The Hobbit, but eventually developed into a much larger work. " +
					"Written in stages between 1937 and 1949, " +
					"The Lord of the Rings is one of the best-selling novels ever written, with over 150 million copies sold.",
				category: []string{"Fantasy", "Adventure"},
			},
			msg: "Book should be updated",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks_controller.NewIBookController(t)
			m.On("Update", tt.args.param, tt.args.title, tt.args.author, tt.args.synopsis, tt.args.category).Return(tt.want)

			got := m.Update(tt.args.param, tt.args.title, tt.args.author, tt.args.synopsis, tt.args.category)
			assertions.Equalf(tt.want, got, tt.msg)
		})
	}
}
