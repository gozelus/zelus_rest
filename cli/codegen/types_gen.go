package codegen

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/gozelus/zelus_rest/cli/tpls"
	"github.com/iancoleman/strcase"
)

type TypeInfo struct {
	TypeName string
	Fields   []*Field
	Comment  string
}
type TypesGenner struct {
	Types []*TypeInfo

	apiFile io.Reader
	pkgName string
	write   io.Writer
}

func NewTypesInfo(writeFile io.Writer, apiFile io.Reader, pkgName string) *TypesGenner {
	return &TypesGenner{apiFile: apiFile, write: writeFile, pkgName: pkgName}
}

func (t *TypesGenner) GenCode() (err error) {
	if err = t.readAllTypeLinesStr(); err != nil {
		return
	}
	if _, err = t.write.Write([]byte("package " + t.pkgName + "\n")); err != nil {
		return err
	}
	var temp *template.Template
	if temp, err = template.New("types gen tpl").Parse(tpls.TypesTpl); err != nil {
		return err
	}
	if err = temp.Execute(t.write, t); err != nil {
		return err
	}
	return nil
}

func (t *TypesGenner) readAllTypeLinesStr() error {
	reader := bufio.NewReader(t.apiFile)
	var lines []string
	var documentStack []string
	var lineNum int
	var lastType = &TypeInfo{}
	var typeDefineBegin bool
	for {
		lineBytes, _, err := reader.ReadLine()
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		if errors.Is(err, io.EOF) {
			return nil
		}
		lineStr := string(lineBytes)
		lines = append(lines, lineStr)
		if strings.HasPrefix(lineStr, "//") {
			documentStack = append(documentStack, lineStr)
			continue
		}
		if strings.HasPrefix(lineStr, "type") {
			// new a field
			typeDefineBegin = true
			lastType.TypeName = strings.Split(lineStr, " ")[1]
			continue
		}
		if strings.HasPrefix(lineStr, "}") && typeDefineBegin {
			// end a type define
			typeDefineBegin = false
			newT := &TypeInfo{
				TypeName: lastType.TypeName,
				Fields:   lastType.Fields,
				Comment:  "",
			}
			t.Types = append(t.Types, newT)
			for i, comment := range documentStack {
				newT.Comment = newT.Comment + comment
				if i != len(documentStack)-1 {
					newT.Comment += "\n"
				}
			}
			documentStack = []string{}
			// empty fields
			lastType.Fields = []*Field{}
			continue
		}
		if len(lineStr) > 0 && typeDefineBegin {
			lineStr = strings.TrimLeft(lineStr, " ")
			keys := strings.Fields(lineStr)
			var f *Field
			if len(keys) == 1 {
				// 内联
				f = &Field{
					Name:         keys[0],
				}
			} else {
				if len(keys) < 3 {
					return errors.New(fmt.Sprintf("field : %s is valid, plz check tag exists", keys[0]))
				}
				f = &Field{
					Name:         keys[0],
					TypeName:     keys[1],
					LowCamelName: strcase.ToLowerCamel(keys[0]),
					Tags:         keys[2],
				}

			}
			if len(keys) >= 4 {
				if len(keys) == 4 {
					f.Comment = keys[3]
				} else {
					f.Comment = "// " + keys[4]
				}
			}
			lastType.Fields = append(lastType.Fields, f)
		}
		lineNum++
	}
}
