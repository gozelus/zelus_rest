package rest

import "net/http"

// 一些预定好的错误
var (
	statusMethodNotAllowed = rsp{
		RCode:    http.StatusMethodNotAllowed,
		RMessage: "method not allowed",
	}
	statusNotFound = rsp{
		RCode:    http.StatusNotFound,
		RMessage: "not found",
	}
	statusInternalServerError = rsp{
		RCode:    http.StatusInternalServerError,
		RMessage: "internal server error",
	}
	statusBadRequest = rsp{
		RCode:    http.StatusBadRequest,
		RMessage: "bad request",
	}
)

type rsp struct {
	RCode    int
	RMessage string
	RData    interface{}
}

var _ Rsp = rsp{}

func (r rsp) ErrorCode() int {
	return r.RCode
}
func (r rsp) ErrorMessage() string {
	return r.RMessage
}
func (r rsp) Data() interface{} {
	return r.RData
}

func StatusMethodNotAllowedResp() Rsp {
	return statusMethodNotAllowed
}
func StatusNotFoundResp() Rsp {
	return statusNotFound
}
func StatusInternalServerErrorResp() Rsp {
	return statusInternalServerError
}
func StatusBadRequestResp() Rsp {
	return statusBadRequest
}
