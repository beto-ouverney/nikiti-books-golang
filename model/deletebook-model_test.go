package model

import (
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/beto-ouverney/nikiti-books/model/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookModel_Delete(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want *customerror.CustomError
		msg  string
	}{
		{
			name: "Should be able to delete a book by title",
			args: args{
				param: "The Lord of the Rings",
			},
			want: nil,
			msg:  "Book should be deleted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := new(mocks.IBookModel)

			m.On("Delete", tt.args.param).Return(nil)

			err := m.Delete(tt.args.param)

			assert.Nil(err, "Error should be nil")
		})
	}
}
