package routes

import (
	"net/http"

	rest "github.com/gozelus/zelus_rest"
	v1 "github.com/gozelus/zelus_rest/example/id_generator/internal/controllers/v1"
	"github.com/gozelus/zelus_rest/example/id_generator/internal/service"
)

var idController = v1.NewIdController(service.NewIdService())
var Routes = []rest.Route{
	{
		Path:    "/v1/id/batch",
		Method:  http.MethodGet,
		Handler: idController.Batch,
	},
	{
		Path:    "/v1/id/single",
		Method:  http.MethodGet,
		Handler: idController.Single,
	},
}
