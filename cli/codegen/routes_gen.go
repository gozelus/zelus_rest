package codegen

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"html/template"
	"io"
	"sort"
)

var routesTpl = `package routes

import (
	"github.com/gozelus/zelus_rest"
	"time"

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
		NeedAuthentication: {{ .NeedAuthentication }},
		AllowCORS: {{ .AllowCORS }},
		{{ if not (eq .TimeoutMs 0) }}TimeOut : {{ .TimeoutMs }} * time.Millisecond,{{ end }}
	},{{ end }}{{ end }}
}
`

type ControllerSort []Controller

func (l ControllerSort) Len() int {
	return len(l)
}

func (l ControllerSort) Less(i, j int) bool {
	return l[i].Name < l[j].Name
}

func (l ControllerSort) Swap(i, j int) {
	l[i], l[j] = l[i], l[j]
}

type RouteGenner struct {
	ModuleName  string
	Controllers ControllerSort
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
	var cs ControllerSort
	for _, c := range controllers {
		cs = append(cs, c)
	}
	sort.Sort(cs)
	r.Controllers = cs
	var t *template.Template
	var err error
	if t, err = template.New("routes new tpl").Funcs(sprig.FuncMap()).Parse(routesTpl); err != nil {
		return nil
	}
	return t.Execute(w, r)
}
