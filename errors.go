package rest

import "net/http"

// 一些预定好的错误
var (
	statusMethodNotAllowed = &statusError{
		RCode:    http.StatusMethodNotAllowed,
		RMessage: "method not allowed",
	}
	statusNotFound = &statusError{
		RCode:    http.StatusNotFound,
		RMessage: "not found",
	}
	statusInternalServerError = &statusError{
		RCode:    http.StatusInternalServerError,
		RMessage: "internal server error",
	}
)

type statusError struct {
	RCode    int
	RMessage string
	RData    interface{}
}

func (s statusError) Error() string {
	return s.RMessage
}

func (s statusError) GetCode() int {
	return s.RCode
}

func (s statusError) GetMessage() string {
	return s.RMessage
}

var _ StatusError = &statusError{}
