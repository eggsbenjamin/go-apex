package lambda

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	tests := []struct {
		typ      string
		err      error
		expected *Error
	}{
		{
			"test type 1",
			errors.New("test msg 1"),
			&Error{
				Type:       "test type 1",
				Message:    "test msg 1",
				StackTrace: []string{"test msg 1", ""},
			},
		},
		{
			"test type 2",
			nil,
			&Error{
				Type:       "test type 2",
				Message:    "",
				StackTrace: nil,
			},
		},
	}

	for i, tc := range tests {
		assert.Equal(t, tc.expected, NewError(tc.typ, tc.err), "error in test case %d", i)
	}
}
