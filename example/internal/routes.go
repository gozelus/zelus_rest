package internal

import (
	"net/http"

	rest "github.com/gozelus/zelus_rest"
)

var (
	routes []rest.Route
)

func init() {
	routes = []rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/user/info",
			Handler: GetUser,
		},
	}
}
