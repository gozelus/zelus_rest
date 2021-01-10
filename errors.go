package rest

import "net/http"

// 一些预定好的错误
var (
	statusUnauthorized = &statusError{
		RCode:    http.StatusUnauthorized,
		RMessage: "status unauthorized",
		Reason: &reason{
			Message: "该接口需要登录",
			Code:    40003,
		},
	}
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

func (s *statusError) WithReason(msg string, code int) StatusError {
	e := statusError{}
	e.Reason = &reason{
		Message: msg,
		Code:    code,
	}
	e.RData = s.RData
	e.RMessage = s.RMessage
	e.RCode = s.RCode
	return &e
}

func (s *statusError) GetReason() interface {
	GetReasonMessage() string
	GetReasonCode() int
} {
	return s.Reason
}

func (s *statusError) Error() string {
	return s.RMessage
}

func (s *statusError) GetCode() int {
	return s.RCode
}

func (s *statusError) GetMessage() string {
	return s.RMessage
}

var _ StatusError = &statusError{}
