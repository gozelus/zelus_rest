package codegen

import (
	"github.com/gozelus/zelus_rest/cli/tpls"
	"github.com/iancoleman/strcase"
	"html/template"
	"io"
	"os"
)

type constructorInfo struct {
	ReturnName   string
	PkgName      string
	PkgNameCamel string
	Imports      []string
}
type WireGenner struct {
	Constructors []*constructorInfo
	Controllers  []*Controller
}

func NewWireGenner(controllers []*Controller) *WireGenner {
	return &WireGenner{
		Controllers: controllers,
	}
}

func (w *WireGenner) initConstructors() error {
	for _, c := range w.Controllers {
		info := &constructorInfo{
			ReturnName:   c.Name,
			PkgName:      c.PkgName,
			PkgNameCamel: strcase.ToCamel(c.PkgName),
		}
		w.Constructors = append(w.Constructors, info)
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
