package codegen

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type handler struct {
	method string
	path   string
	name   string
}
type ControllerGenner struct {
	readerFiler io.Reader
	handlers    []*handler
	reader      io.Reader
	writer      io.Writer
}

func NewControllerGenner(file io.Reader, writer io.Writer) *ControllerGenner {
	return &ControllerGenner{writer: writer, reader: file}
}

func (c *ControllerGenner) GenCode() error {
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
		lineNum ++
		lineStr := string(lineBytes)
		lineStrs = append(lineStrs, lineStr)
		if strings.HasPrefix(lineStr, "service") {
			// service define begin
			serviceDefineNum = lineNum
			defineServiceBegin = true
			continue
		}
		if strings.HasPrefix(lineStr, "}") && defineServiceBegin {
			// service define end
			serviceEndDefineNum = lineNum
			defineServiceBegin = false
			serviceLines = append(serviceLines, []int{serviceDefineNum, serviceEndDefineNum})
			continue
		}
	}
	for _, define := range serviceLines {
		fmt.Println(lineStrs[define[0]:define[1]])
	}
	return nil
}
