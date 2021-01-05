package codegen

import (
	"fmt"
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
	Imports      []string
}

type ServiceGenner struct {
	serviceInfo *serviceInfo
	moduleName  string
}

func NewServiceGener(moduleName string) *ServiceGenner {
	s := &ServiceGenner{moduleName: moduleName}
	return s
}

// file 要写入的文件
// controller 要服务的 controller
func (s *ServiceGenner) GenCode(file io.Writer, c *Controller) error {
	s.serviceInfo = s.initServiceInfo(c)
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

func (s *ServiceGenner) initServiceInfo(controller *Controller) *serviceInfo {
	return &serviceInfo{
		Name:     controller.Name,
		Handlers: controller.Handlers,
		Imports: []string{
			fmt.Sprintf(`"%s/internal"`, s.moduleName),
			fmt.Sprintf(`"%s/api/errors"`, s.moduleName),
		},
		PkgName:      strings.Split(controller.PkgName, "_")[0] + "_services",
		TypesPkgName: controller.TypesPkgName,
	}
}
