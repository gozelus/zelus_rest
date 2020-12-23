package main

var RouteTpl = `package router

import (
	"github.com/gozelus/zelus_rest"
)

{{range .Controllers}}
type {{.Name}}ControllerInter interface {
	{{range .Handlers}} {{.FuncName}}(w http.ResponseWriter, req *http.Request)
	{{end}}
}
{{end}}

type Router struct {
	{{range .Controllers}} {{.Name}} {{.Name}}ControllerInter {{end}}
}

func NewRouter({{range .Controllers}}
{{.Name}} {{.Name}}ControllerInter,{{end}}) *Router {
	return &Router{ {{range.Controllers}}
		{{.Name}}: {{.Name}}, {{end}}
	}
}

func (r *Router) Routes() []rest.Route {
	return []rest.Route{
		{{range .Controllers}}
		{{range .Handlers}}
		{
			Method:  "{{.Method}}",
			Path:    "{{.Path}}",
			Handler: r.{{.ControllerName}}.{{.FuncName}},
		},
        {{end}}
		{{end}}
	}
}
`
