package tpls

var ControllerTpl = `package {{ .PkgName }}

import (
	rest "github.com/gozelus/zelus_rest"
	{{ range $import := .Imports }}
	{{ $import }}
	{{ end }}
)


type {{ .Name }}Controller struct {
	service *{{ $.ServicesPkgName }}.{{ .Name }}Service
}
func New{{ .Name }}Controller(service *{{ $.ServicesPkgName }}.{{ .Name }}Service) *{{ .Name }}Controller {
	return &{{ .Name }}Controller{service : service}
}

{{ range .Handlers }}
func (c *{{ $.Name }}Controller) {{ .Name }}(ctx rest.Context) {
	res := &{{ $.TypesPkgName }}.{{ .ResponseType }}{}
	req := &{{ $.TypesPkgName }}.{{ .RequestType }}{}
	var err error 
	if err := ctx.{{if ne .Method "GET" }}JSONBodyBind{{ else }}JSONQueryBind{{ end }}(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason("部分参数校验不通过", 40000))
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
