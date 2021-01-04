package tpls

var ControllerTpl = `package {{ .PkgName }}
type {{ .Name }}Service interface {
}
type {{ .Name }}Controller struct {
	service {{ .Name }}Service
}
func New{{ .Name }}Controller(service {{ .Name }}Service) *{{ .Name }}Controller {
	return &{{ .Name }}Controller{service : service}
}
`
