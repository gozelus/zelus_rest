package codegen

import (
	"github.com/gozelus/zelus_rest/cli/tpls"
	"html/template"
	"io"
	"strings"
)

type serviceInfo struct {
	Name         string
	Handlers     []*handler
	PkgName      string
	TypesPkgName string
}

type ServiceGenner struct {
	controller  *controller
	serviceInfo *serviceInfo
}

func NewServiceGener(c *controller) *ServiceGenner {
	s := &ServiceGenner{
		controller:  c,
		serviceInfo: initServiceInfo(c),
	}
	return s
}

// file 要写入的文件
// controller 要服务的 controller
func (s *ServiceGenner) GenCode(file io.Writer) error {
	var t *template.Template
	var err error
	if t, err = template.New("service new").Parse(tpls.ServiceTpl); err != nil {
		return err
	}
	if err := t.Execute(file, s.serviceInfo); err != nil {
		return err
	}
	return nil
}

func initServiceInfo(controller *controller) *serviceInfo {
	return &serviceInfo{
		Name:         controller.Name,
		Handlers:     controller.Handlers,
		PkgName:      strings.Split(controller.PkgName, "_")[0] + "_services",
		TypesPkgName: controller.TypesPkgName,
	}
}
