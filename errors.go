package rest

import "net/http"

// 一些预定好的错误
var (
	statusMethodNotAllowed = &statusError{
		RCode:    http.StatusMethodNotAllowed,
		RMessage: "method not allowed",
		Reason: &reason{
			Message: "找不到对应的方法",
			Code:    40002,
		},
	}
	statusNotFound = &statusError{
		RCode:    http.StatusNotFound,
		RMessage: "not found",
		Reason: &reason{
			Message: "找不到对应的路径",
			Code:    40001,
		},
	}
	statusInternalServerError = &statusError{
		RCode:    http.StatusInternalServerError,
		RMessage: "internal server error",
		Reason: &reason{
			Message: "服务器好像开小差了",
			Code:    50000,
		},
	}
)

type reason struct {
	Message string
	Code    int
}

func (r *reason) GetReasonMessage() string {
	return r.Message
}

func (r *reason) GetReasonCode() int {
	return r.Code
}

type statusError struct {
	RCode    int
	RMessage string
	RData    interface{}
	Reason   *reason
}

func (s statusError) GetReason() interface {
	GetReasonMessage() string
	GetReasonCode() int
} {
	return s.Reason
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
