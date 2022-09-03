package test

import (
	mocks_controller "github.com/beto-ouverney/nikiti-books/controller/mocks"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookController_FindAll(t *testing.T) {
	assertions := assert.New(t)

	tests := []struct {
		name  string
		want  []byte
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name:  "Should be able to find all books",
			want:  []byte(`[{"title":"The Lord of the Rings","author":"J.R.R. Tolkien","category":["Fantasy","Adventure"],"synopsis":"The Lord of the Rings is an epic high-fantasy novel by English author and scholar J. R. R. Tolkien."},{"title":"The Hobbit","author":"J.R.R. Tolkien","category":["Fantasy","Adventure"],"synopsis":"The Hobbit is a children's fantasy novel by English author J. R. R. Tolkien."}]`),
			want1: nil,
			msg:   "Books should be found",
			msg1:  "Error should be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks_controller.NewIBookController(t)
			m.On("FindAll").Return(tt.want, tt.want1)

			got, got1 := m.FindAll()
			assertions.Equalf(tt.want, got, tt.msg)
			assertions.Equalf(tt.want1, got1, tt.msg1)
		})
	}
}
