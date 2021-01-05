package codegen

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"io"
	"html/template"
)

var routesTpl = `package routes

import (
	"github.com/gozelus/zelus_rest"

	{{ range $import, $key := .Imports }}
	"{{ $import }}"
	{{ end }}
)

{{ range $controller := .Controllers }}
var {{ lower .Name }}Controller = injector.New{{ camelcase $controller.PkgName }}{{ $controller.Name }}Controller() {{ end }}
var Routes = []rest.Route {	{{ range $controller := .Controllers }} {{ range .Handlers }}
	{{ range $comment := .Comments }}
	{{ $comment }} {{ end }}
	{
		Method:  "{{ .Method }}",
		Path:    "{{ .Path }}",
		Handler: {{ lower $controller.Name }}Controller.{{ .Name }},
	},{{ end }}{{ end }}
}
`

type RouteGenner struct {
	ModuleName  string
	Controllers []Controller
	Imports     map[string]string
}

func NewRouteGenner(moduleName string) *RouteGenner {
	r := &RouteGenner{ModuleName: moduleName}
	r.Imports = map[string]string{
		fmt.Sprintf(`%s/internal/injector`, moduleName): fmt.Sprintf(`"%s/internal/injector"`, moduleName),
	}
	return r
}

func (r *RouteGenner) GenCode(w io.Writer, controllers []Controller) error {
	r.Controllers = controllers
	var t *template.Template
	var err error
	if t, err = template.New("routes new tpl").Funcs(sprig.FuncMap()).Parse(routesTpl); err != nil {
		return nil
	}
	return t.Execute(w, r)
}
