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
	"strings"
)

type controller struct {
	Name     string
	Handlers []*handler
	PkgName  string
	// 依赖的类型包名
	TypesPkgName string
	// 依赖的服务包名
	ServicesPkgName string
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

	// reader api 文件的读取入口抽象
	reader io.Reader
	// 依赖的类型包名
	TypesPkgName string
	// 依赖的服务包名
	ServicesPkgName string

	// key1 一级path  ->  文件夹名
	// key2 二级path  ->  文件名
	// 如 /v1/user/create -> { "v1" : { "user" : $controller } }
	// controller 下的函数名，将会被 @handler 后的字符串映射
	Group map[string]map[string]*controller
}

func NewControllerGenner() *ControllerGenner {
	return &ControllerGenner{Group: map[string]map[string]*controller{}}
}

// 将 controller 结构体转为代码写入文件
func (c ControllerGenner) GenCode(w io.Writer, controller *controller) error {
	if err := c.execTemplate(w, controller); err != nil {
		return err
	}
	return nil
}

// 通过 api 文件生成 controller 定义的结构体
func (c *ControllerGenner) ParseApiFile(file io.Reader, typesPkgName string) (map[string]map[string]*controller, error) {
	c.reader = file
	c.TypesPkgName = typesPkgName
	if err := c.initHandlers(); err != nil {
		return nil, err
	}
	return c.Group, nil
}

func (c *ControllerGenner) execTemplate(w io.Writer, controller *controller) error {
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
					Name:            strcase.ToCamel(controllerName),
					Handlers:        []*handler{h},
					PkgName:         group + "_controllers",
					TypesPkgName:    c.TypesPkgName,
					ServicesPkgName: group + "_services",
				}
			}
		} else {
			c.Group[group] = map[string]*controller{
				controllerName: {
					Name:            strcase.ToCamel(controllerName),
					Handlers:        []*handler{h},
					PkgName:         group + "_controllers",
					TypesPkgName:    c.TypesPkgName,
					ServicesPkgName: group + "_services",
				},
			}
		}
		h = nil
	}
	return nil
}
