package http

import "net/http"

var (
	ErrBadRequestBodyMustBeJson = NewError(http.StatusBadRequest, "request body must be json")
)

type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(statusCode int, message string) *Error {
	return &Error{
		StatusCode: statusCode,
		Message:    message,
	}
}
