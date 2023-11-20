package errs

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", ce.Code, ce.Message)
}

func New(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
