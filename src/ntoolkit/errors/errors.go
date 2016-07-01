package errors

import "fmt"

// Error is a basic common error type that wraps a message, a code and
// an optional inner error.
type Error struct {
	Reason string
	Inner  error
	Code   int
}

// Error interface
func (err Error) Error() string {
	if err.Inner == nil {
		return fmt.Sprintf("(%d) %s", err.Code, err.Reason)
	}
	return fmt.Sprintf("(%d) %s: %s", err.Code, err.Reason, err.Inner)
}

// Fail returns an error with a message
func Fail(code int, inner error, format string, args ...interface{}) error {
	return Error{
		fmt.Sprintf(format, args...),
		inner,
		code}
}

// Is checks for the error code on a standard error
func Is(err error, code int) bool {
	defer func() {
		fmt.Println("recovered:", recover())
	}()
	return code == err.(Error).Code
}
