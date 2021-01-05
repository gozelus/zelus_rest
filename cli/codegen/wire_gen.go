package codegen

import (
	"fmt"
	"github.com/gozelus/zelus_rest/cli/tpls"
	"github.com/iancoleman/strcase"
	"text/template"
	"io"
	"os"
	"strings"
)

type constructorInfo struct {
	ReturnName   string
	PkgName      string
	PkgNameCamel string
}
type WireGenner struct {
	Constructors []*constructorInfo
	Controllers  []Controller
	ModuleName   string
	Imports      map[string]string
}

func NewWireGenner(controllers []Controller, moduleName string) *WireGenner {
	return &WireGenner{
		ModuleName:  moduleName,
		Controllers: controllers,
		Imports: map[string]string{},
	}
}

func (w *WireGenner) initConstructors() error {
	for _, c := range w.Controllers {
		controllerInfo := &constructorInfo{
			ReturnName:   c.Name + "Controller",
			PkgName:      c.PkgName,
			PkgNameCamel: strcase.ToCamel(c.PkgName),
		}
		c.PkgName = strings.ReplaceAll(c.PkgName, "controllers", "services")
		serviceInfo := &constructorInfo{
			ReturnName:   c.Name + "Service",
			PkgName:      c.PkgName,
			PkgNameCamel: strcase.ToCamel(c.PkgName),
		}
		w.Constructors = append(w.Constructors, controllerInfo, serviceInfo)
		key1 := fmt.Sprintf(`"%s/internal/controllers/%s"`, w.ModuleName, c.Group)
		key2 := fmt.Sprintf(`"%s/internal/services/%s"`, w.ModuleName, c.Group)
		w.Imports[key1] = key1
		w.Imports[key2] = key2
	}
	return nil
}

func (w *WireGenner) GenCode(zelusCtlWireFile, selfWireFile *os.File) error {
	if err := w.initConstructors(); err != nil {
		return err
	}
	var codegen = func(file io.Writer, tpl string) error {
		var t *template.Template
		var err error
		if t, err = template.New("wire zelus ctl gen").Parse(tpl); err != nil {
			return err
		}
		if err = t.Execute(file, w); err != nil {
			return err
		}
		return nil
	}
	if selfWireFile != nil {
		if err := codegen(selfWireFile, tpls.WireSelfTpls); err != nil {
			return err
		}
	}
	if zelusCtlWireFile != nil {
		if err := codegen(zelusCtlWireFile, tpls.WireZelusTpls); err != nil {
			return err
		}
	}
	return nil
}
