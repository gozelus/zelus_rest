package codegen

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/gozelus/zelus_rest/cli/tpls"
	"github.com/iancoleman/strcase"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type controller struct {
	Name         string
	Handlers     []*handler
	PkgName      string
	TypesPkgName string
}
type handler struct {
	Method       string
	Path         string
	Name         string
	RequestType  string
	ResponseType string
	Comments     []string
}
type ControllerGenner struct {
	readerFiler  io.Reader
	reader       io.Reader
	writer       io.Writer
	genPath      string
	typesPkgName string

	Group map[string]map[string]*controller
}

func NewControllerGenner(file io.Reader, genPath, typesPkgName string) *ControllerGenner {
	return &ControllerGenner{typesPkgName: typesPkgName, genPath: genPath, reader: file, Group: map[string]map[string]*controller{}}
}

func (c *ControllerGenner) GenCode() error {
	if err := c.initHandlers(); err != nil {
		return err
	}
	if err := c.initDir(); err != nil {
		return err
	}
	return nil
}

func (c *ControllerGenner) initDir() error {
	for group, controllerMap := range c.Group {
		path := filepath.Join(c.genPath, group)
		_, err := os.Stat(path)
		if err == nil {
			fmt.Println(color.GreenString("%s exist, will remove and recreate", path))
			if err := os.RemoveAll(path); err != nil {
				return err
			}
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
		}
		if os.IsNotExist(err) {
			fmt.Println(color.GreenString("%s not exist, will recreate", path))
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
		}
		for _, controller := range controllerMap {
			filename := filepath.Join(path, strcase.ToSnake(controller.Name+"_controller.go"))
			w, err := os.Create(filename)
			if err != nil {
				return err
			}
			fmt.Println(color.GreenString("%s created", filename))
			if err := c.execTemplate(w, controller, group); err != nil {
				return err
			}
		}
	}
	return nil
}
func (c *ControllerGenner) execTemplate(w io.Writer, controller *controller, groupName string) error {
	var t *template.Template
	var err error
	if t, err = template.New("controller new").Parse(tpls.ControllerTpl); err != nil {
		return err
	}
	return t.Execute(w, controller)
}
func (c *ControllerGenner) initHandlers() error {
	reader := bufio.NewReader(c.reader)
	var lineNum int
	var serviceLines [][]int
	var serviceDefineNum int
	var serviceEndDefineNum int
	var lineStrs []string
	var defineServiceBegin bool
	for {
		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		lineStr := string(lineBytes)
		lineStr = strings.TrimSpace(lineStr)
		if len(lineStr) == 0 {
			continue
		}
		lineStrs = append(lineStrs, lineStr)
		lineNum++
		if strings.HasPrefix(lineStr, "service") {
			// service define begin
			serviceDefineNum = lineNum
			defineServiceBegin = true
			continue
		}
		if strings.HasPrefix(lineStr, "}") && defineServiceBegin {
			// service define end
			serviceEndDefineNum = lineNum - 1
			defineServiceBegin = false
			serviceLines = append(serviceLines, []int{serviceDefineNum, serviceEndDefineNum})
			continue
		}
	}
	for _, define := range serviceLines {
		if err := c.handleHandlerLine(lineStrs[define[0]:define[1]]); err != nil {
			return err
		}
	}
	return nil
}
func (c *ControllerGenner) handleHandlerLine(lines []string) error {
	var h *handler
	for _, line := range lines {
		fmt.Println(color.BlueString("line : %s", line))
		line = strings.TrimLeft(line, " ")
		keys := strings.Split(line, " ")
		if strings.HasPrefix(line, "//") {
			h = &handler{}
			h.Comments = append(h.Comments, line)
			continue
		}
		if strings.HasPrefix(line, "@") {
			if len(keys) < 2 {
				return errors.New(fmt.Sprintf("line : %s is valid, check if u have the handler name", line))
			}
			if h == nil {
				h = &handler{}
			}
			h.Name = strcase.ToCamel(keys[1])
			continue
		}
		if len(keys) != 5 {
			return errors.New(fmt.Sprintf("line : %s is valid, check if u have req and res", line))
		}
		h.Method = strings.ToUpper(keys[0])
		h.Path = keys[1]
		h.RequestType = keys[2]
		h.ResponseType = keys[4]

		// 按照一级path，认为group
		path := strings.Split(h.Path, "/")
		if len(path) < 3 {
			return errors.New(fmt.Sprintf("line : %s path : %s is too short, min lenth is 3", line, h.Path))
		}
		group := path[1]
		controllerName := path[2]
		if _, ok := c.Group[group]; ok {
			if excontroller, ok2 := c.Group[group][controllerName]; ok2 {
				excontroller.Handlers = append(excontroller.Handlers, h)
			} else {
				c.Group[group][controllerName] = &controller{
					Name:         strcase.ToCamel(controllerName),
					Handlers:     []*handler{h},
					PkgName:      group,
					TypesPkgName: c.typesPkgName,
				}
			}
		} else {
			c.Group[group] = map[string]*controller{
				controllerName: {
					Name:         strcase.ToCamel(controllerName),
					Handlers:     []*handler{h},
					PkgName:      group,
					TypesPkgName: c.typesPkgName,
				},
			}
		}
		h = nil
	}
	return nil
}
