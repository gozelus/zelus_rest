package controller

import (
	"testing"

	"github.com/gozelus/zelus_rest/cli/types"
)

func TestUser_checkReq(t *testing.T) {
	c := &User{
		UserService: nil,
	}
	c.checkReq(types.UserGetRequest{})
}