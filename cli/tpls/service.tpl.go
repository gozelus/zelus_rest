package tpls

var ServiceTpl = `package {{ .PkgName }}
type {{ .Name }}Service struct {
	// 以后放入要依赖度的对象
}
func New{{ .Name }}Service() *{{ .Name }}Service {
	return &{{ .Name }}Service{}
}	
{{ range .Handlers }}
func (s *{{ $.Name }}Service) {{ .Name }}(ctx rest.Context, request *{{ $.TypesPkgName }}.{{ .RequestType }}) (*{{ $.TypesPkgName }}.{{ .ResponseType }}, error) {
	return nil, errors.New("no imp")
}
{{ end }}
`
