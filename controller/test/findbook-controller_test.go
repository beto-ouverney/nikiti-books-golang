package test

import (
	"encoding/json"
	mocks_controller "github.com/beto-ouverney/nikiti-books/controller/mocks"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookController_FindBook(t *testing.T) {
	assertions := assert.New(t)

	res := entity.Book{
		Title:    "The Lord of the Rings",
		Author:   "J. R. R. Tolkien",
		Category: []string{"Fantasy", "Adventure"},
		Synopsis: "The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien.",
	}

	resJ, errJ := json.MarshalIndent(res, "", "    ")
	if errJ != nil {
		t.Errorf("Error marshalling book: %v", errJ)
	}

	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name: "Should be able to find a book by title",
			args: args{
				name: "The Lord of the Rings",
			},
			want:  resJ,
			want1: nil,
			msg:   "Book should be found",
			msg1:  "Error should be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks_controller.NewIBookController(t)
			m.On("FindBook", tt.args.name).Return(tt.want, nil)

			got, got1 := m.FindBook(tt.args.name)
			assertions.Equalf(tt.want, got, tt.msg)
			assertions.Equalf(tt.want1, got1, tt.msg1)
		})
	}
}
