package test

import (
	mocks_controller "github.com/beto-ouverney/nikiti-books/controller/mocks"
	"github.com/beto-ouverney/nikiti-books/customerror"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookController_Delete(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		name string
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
				name: "The Lord of the Rings",
			},
			want: nil,
			msg:  "Book should be deleted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks_controller.NewIBookController(t)
			m.On("Delete", tt.args.name).Return(tt.want)

			got := m.Delete(tt.args.name)
			assertions.Equalf(tt.want, got, tt.msg)
		})
	}
}
