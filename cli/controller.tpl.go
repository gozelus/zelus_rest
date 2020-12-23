package main

var ControllerTpl = `package controller

import (
	"net/http"

	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/cli/types"
)

type {{.ControllerName}}ServiceInterface interface { {{range .Handlers}}
	{{.FuncName}}(*types.{{.ParamsName}}) (*types.{{.ReturnName}}, error) {{end}}
}

type {{.ControllerName}} struct {
	{{.ControllerName}}Service {{.ControllerName}}ServiceInterface
}

func New{{.ControllerName}}({{.ControllerName}}Service {{.ControllerName}}ServiceInterface) *{{.ControllerName}} {
	return &{{.ControllerName}}{
		{{.ControllerName}}Service: {{.ControllerName}}Service,
	}
}
{{range .Handlers}}
func (c *{{$.ControllerName}}) {{.FuncName}}(w http.ResponseWriter, req *http.Request) {
	param := &types.{{.ParamsName}}{}
    rest.JsonBodyFromRequest(req, param)
	// check with tag
	res, _ := c.{{$.ControllerName}}Service.{{.FuncName}}(param)
	rest.OkJson(w, res)
}
{{end}}
`
