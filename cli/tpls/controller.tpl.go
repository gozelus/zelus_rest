package tpls

var ControllerTpl = `package {{ .PkgName }}

import (
	"github.com/gozelus/zelus_rest"
)

type {{ .Name }}Service interface { {{ range .Handlers }}
	{{ .Name }}(ctx rest.Context, req *{{ $.TypesPkgName }}.{{ .RequestType }}) (*{{ $.TypesPkgName }}.{{ .ResponseType }}, error) {{ end }}
}
type {{ .Name }}Controller struct {
	service {{ .Name }}Service
}
func New{{ .Name }}Controller(service {{ .Name }}Service) *{{ .Name }}Controller {
	return &{{ .Name }}Controller{service : service}
}

{{ range .Handlers }}
func (c *{{ $.Name }}Controller) {{ .Name }}(ctx rest.Context) {
	res := &{{ $.TypesPkgName }}.{{ .ResponseType }}{}
	req := &{{ $.TypesPkgName }}.{{ .RequestType }}{}
	var err error 
	if err := ctx.{{if ne .Method "GET" }}JSONBodyBind{{ else }}JSONQueryBind{{ end }}(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
		return
	}
	if res, err = c.service.{{ .Name }}(ctx, req);err!=nil{
		ctx.RenderErrorJSON(nil, err)
		return
	}
	ctx.RenderOkJSON(res)
}
{{ end }}
`
