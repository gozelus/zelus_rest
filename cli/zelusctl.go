package main

import (
	"bufio"
	"bytes"
	"fmt"
	goformat "go/format"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type Member struct {
	Name string
	Type string
}
type Type struct {
	Name     string
	TypeName string
	Tags     []string
	Members  []Member
}

type Handler struct {
	FuncName       string
	ReturnName     string
	ParamsName     string
	Method         string
	Path           string
	ControllerName string
}
type Controller struct {
	Name     string
	Handlers []Handler
}

// ApiStruct 用于描述 api 服务
type ApiStruct struct {
	Info string
	// 定义的一些结构体
	Types []string
	// api 服务
	Controller Controller
	Imports    string
	// api 服务定义的开始
	serviceBeginLine int
}

func GenGoCode() {
	api, _ := ParseApi("/Users/zhengli/workspace/private/projects/zelus_rest/example/api/minitaobao.api")
	genTypesCode(api)
	genServiceCode(api)
	if err := genRouteCode(api); err != nil {
		panic(err)
	}
}

func genRouteCode(api *ApiStruct) error {
	os.RemoveAll("./routes")
	os.Mkdir("./routes", 0777)
	file, _ := os.Create("./routes/routes.go")
	// 编译模板
	return template.Must(template.New("routesTpl").Parse(RouteTpl)).Execute(file, struct {
		Controllers []Controller
	}{
		Controllers: []Controller{api.Controller},
	})
}
func genTypesCode(api *ApiStruct) error {
	os.RemoveAll("./types")
	os.Mkdir("./types", 0777)
	file, _ := os.Create("./types/type.go")
	code := "package types\n\n\n"
	for _, typeStr := range api.Types {
		code += typeStr
	}
	formatCode, _ := goformat.Source([]byte(code))
	file.Write(formatCode)
	return nil
}
func genServiceCode(api *ApiStruct) error {
	os.RemoveAll("./controllers")
	os.Mkdir("./controllers", 0777)
	goFileName := fmt.Sprintf("./controllers/%s_controller.go", strings.ToLower(api.Controller.Name))
	file, _ := os.Create(goFileName)
	// 编译模板
	controllerName := api.Controller.Name
	return template.Must(template.New("controllerTpl").Parse(ControllerTpl)).Execute(file, struct {
		ControllerName string
		Handlers       []Handler
	}{
		Handlers:       api.Controller.Handlers,
		ControllerName: controllerName,
	})
}

// ParseApi 用于将 api 文件转换为 ApiStruct
// src 文本字符串
func ParseApi(filename string) (*ApiStruct, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	srcBytes, _ := ioutil.ReadAll(file)
	src := string(srcBytes)
	buffer := new(bytes.Buffer)
	buffer.WriteString(src)
	lineNumber := 0
	apiFile := baseState{
		r:          bufio.NewReader(buffer),
		lineNumber: &lineNumber,
	}
	apiStruct := new(ApiStruct)
	return apiFile.process(apiStruct)
}

// baseState 用于存储读取过程
type baseState struct {
	r          *bufio.Reader
	lineNumber *int
}

func (s *baseState) process(api *ApiStruct) (*ApiStruct, error) {
	var builder strings.Builder
	for {
		// 累计字符串
		ch, _ := s.read()

		if ch == 0 {
			token := builder.String()
			fmt.Println(token)
			break
		}
		if isNewline(ch) || isSpace(ch) {
			token := builder.String()
			if token == "type" {
				if err := s.processType(api, token+string(ch)); err != nil {
					return nil, err
				}
			}
			if token == "service" {
				if err := s.processService(api, token+string(ch)); err != nil {
					return nil, err
				}
			}
			builder.Reset()
		} else {
			builder.WriteRune(ch)
		}
	}
	return api, nil
}

// processService 处理 service 开头的代码块
// 此函数会一直读取到 大括号 为止
func (s *baseState) processService(api *ApiStruct, token string) error {
	var controllerName string
	var handlers []Handler

	var innerProcess = func(lines []string) {
		var lineNum int
		for {
			if lineNum > len(lines)-1 {
				return
			}
			line := lines[lineNum]
			if lineNum == 0 {
				controllerName = strings.Trim(strings.Trim(line, "{"), " ")
				lineNum += 1
				continue
			}

			if !strings.Contains(line, "@handler") {
				lineNum += 1
				continue
			}
			// 每隔2行做为一个handler的申明
			handlerName := strings.Trim(strings.Split(line, "@handler")[1], " ")
			line = strings.Trim(lines[lineNum+1], " ")
			statement := strings.Split(line, " ")
			handlers = append(handlers, Handler{
				ControllerName: controllerName,
				FuncName:       handlerName,
				Method:         strings.ToUpper(statement[0]),
				Path:           statement[1],
				ParamsName:     statement[2],
				ReturnName:     statement[4],
			})
			lineNum += 2
		}
	}

	var serviceToken = token
	var err error
	var lineNum int
	var lines []string
	for {
		var next string
		lineNum += 1
		if next, err = s.readLine(); err != nil {
			return err
		}
		serviceToken += next + "\n"
		lines = append(lines, next)
		if next == "}" {
			break
		}
	}
	innerProcess(lines)
	api.Controller = Controller{
		Name: controllerName,
	}
	for _, h := range handlers {
		api.Controller.Handlers = append(api.Controller.Handlers, h)
	}
	for _, h := range api.Controller.Handlers {
		fmt.Printf("method : %s \n", h.Method)
		fmt.Printf("path : %s \n", h.Path)
		fmt.Printf("req : %s \n", h.ParamsName)
		fmt.Printf("res : %s \n", h.ReturnName)
		fmt.Printf("name : %s \n", h.FuncName)
	}
	// service 开头后 读取 空格-> { 之间的字符串做为 controllerName
	//controllerName := ""
	//// 找到 @ 符号开头的行，然后取当下行的下 2 行做为一个controller的接口
	//api.Service = serviceToken
	return nil
}

// processType 处理 type 开头的代码块
// 此函数会一直读取到 大括号 为止
func (s *baseState) processType(api *ApiStruct, token string) error {
	var structToken = token
	var err error
	for {
		var next string
		if next, err = s.readLine(); err != nil {
			return err
		}
		structToken += next + "\n"
		if next == "}" {
			break
		}
	}
	api.Types = append(api.Types, structToken)
	return nil
}

// readLine 读取完整的一行
func (s *baseState) readLine() (string, error) {
	line, _, err := s.r.ReadLine()
	if err != nil {
		return "", err
	}
	*s.lineNumber++
	return string(line), nil
}
func (s *baseState) read() (rune, error) {
	value, err := read(s.r)
	if err != nil {
		return 0, err
	}
	if isNewline(value) {
		*s.lineNumber++
	}
	return value, nil
}
func isNewline(r rune) bool {
	return r == '\n' || r == '\r'
}
func read(r *bufio.Reader) (rune, error) {
	ch, _, err := r.ReadRune()
	return ch, err
}
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}
