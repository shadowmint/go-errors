package errors

import (
	"fmt"
	"reflect"
)

// Error is a basic common error type that wraps a message, a code and
// an optional inner error.
type Error struct {
	Reason string
	Inner  error
	Type   reflect.Type
}

// Error interface
func (err Error) Error() string {
	if err.Inner == nil {
		return fmt.Sprintf("(%v) %s", err.Type, err.Reason)
	}
	return fmt.Sprintf("(%v) %s: %s", err.Type, err.Reason, err.Inner)
}

// Fail returns an error with a message.
func Fail(etype interface{}, inner error, format string, args ...interface{}) error {
	return Error{
		fmt.Sprintf(format, args...),
		inner,
		reflect.TypeOf(etype),
	}
}

// Is checks for the error type on a standard error
func Is(err error, etype interface{}) bool {
	stderr, ok := err.(Error)
	if ok {
		return stderr.Type == reflect.TypeOf(etype)
	}
	return false
}
