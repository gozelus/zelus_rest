package main

import (
	"bufio"
	"bytes"
	"fmt"
	goformat "go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Member struct {
	Name string
	Type string
}
type Type struct {
	Name    string
	Members []Member
}

// ApiStruct 用于描述 api 服务
type ApiStruct struct {
	Info string
	// 定义的一些结构体
	Types []string
	// api 服务
	Service string
	Imports string
	// api 服务定义的开始行
	serviceBeginLine int
}

func GenGoCode() {
	api, _ := ParseApi("/Users/zhengli/workspace/projects/zelus_rest/example/api/minitaobao.api")

	if _, err := os.Stat("./types"); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("types", 0777)
		} else {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat("./types/types.go"); err != nil {
		if os.IsNotExist(err) {
			file, _ := os.Create("./types/type.go")
			code := "package types\n\n\n"
			for _, typeStr := range api.Types {
				code += typeStr
			}
			formatCode, _ := goformat.Source([]byte(code))
			file.Write(formatCode)
		} else {
			log.Fatal(err)
		}
	}
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
			builder.Reset()
		} else {
			builder.WriteRune(ch)
		}
	}
	return api, nil
}

// processType 处理 type 开头的代码块
// 此函数会一直读取到 大括号为止
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
