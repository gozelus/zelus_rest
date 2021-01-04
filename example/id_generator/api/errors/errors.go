package apiErrors

import (
	"fmt"

	rest "github.com/gozelus/zelus_rest"
)

type StatusError struct {
	Err     error
	Message string
	Code    int
	Reason  string
}

func (s *StatusError) Error() string {
	return s.Message
}

func (s *StatusError) GetCode() int {
	return s.Code
}

func (s *StatusError) GetMessage() string {
	if len(s.Reason) > 0 {
		return fmt.Sprintf("%s for %s", s.Message, s.Reason)
	}
	return s.Message
}

func (s *StatusError) WithReason(reason string) *StatusError {
	s2 := &StatusError{}
	s2.Message = s.Message
	s2.Err = s.Err
	s2.Message = s.Message
	s2.Code = s.Code
	s2.Reason = reason
	return s2
}

var _ rest.StatusError = &StatusError{}
