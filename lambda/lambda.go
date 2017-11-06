package lambda

import (
	"bytes"
	"fmt"
	"strings"
)

// Error represents the expected node.js lambda error format.
type Error struct {
	Message    string   `json:"errorMessage"`
	Type       string   `json:"errorType"`
	StackTrace []string `json:"stackTrace"`
}

// Error implements error.
func (e Error) Error() string {
	return e.Message
}

// NewError is the constructor for Error.
func NewError(typ string, err error) *Error {
	var (
		msg string
		st  []string
	)

	if err != nil {
		msg = err.Error()
		st = createStackTrace(err)
	}

	return &Error{
		Type:       typ,
		Message:    msg,
		StackTrace: st,
	}
}

// createStackTrace creates an array of strings from an error's stacktrace.
func createStackTrace(err error) []string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%+v\n", err)
	return strings.Split(b.String(), "\n")
}
