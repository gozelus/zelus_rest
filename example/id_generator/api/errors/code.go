package apiErrors

import "fmt"

var BadRequest = &StatusError{
	Err:     fmt.Errorf("bad request"),
	Message: "BadRequest",
	Code:    400,
}
