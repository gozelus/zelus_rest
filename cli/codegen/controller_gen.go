package codegen

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type controller struct {
	Name     string
	Handlers []*handler
}
type handler struct {
	method       string
	path         string
	name         string
	requestType  string
	responseType string
	comments     []string
}
type ControllerGenner struct {
	readerFiler io.Reader
	Group       map[string]map[string]*controller
	reader      io.Reader
	writer      io.Writer
}

func NewControllerGenner(file io.Reader, writer io.Writer) *ControllerGenner {
	return &ControllerGenner{writer: writer, reader: file, Group: map[string]map[string]*controller{}}
}

func (c *ControllerGenner) GenCode() error {
	if err := c.initHandlers(); err != nil {
		return err
	}
	return nil
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
		fmt.Printf("line : %s \n", line)
		line = strings.TrimLeft(line, " ")
		keys := strings.Split(line, " ")
		if strings.HasPrefix(line, "//") {
			h = &handler{}
			h.comments = append(h.comments, line)
			continue
		}
		if strings.HasPrefix(line, "@") {
			if len(keys) < 2 {
				return errors.New(fmt.Sprintf("line : %s is valid, check if u have the handler name", line))
			}
			if h == nil {
				h = &handler{}
			}
			h.name = keys[1]
			continue
		}
		if len(keys) != 5 {
			return errors.New(fmt.Sprintf("line : %s is valid, check if u have req and res", line))
		}
		h.method = keys[0]
		h.path = keys[1]
		h.requestType = keys[2]
		h.responseType = keys[4]

		// 按照一级path，认为group
		path := strings.Split(h.path, "/")
		if len(path) < 3 {
			return errors.New(fmt.Sprintf("line : %s path : %s is too short, min lenth is 3", line, h.path))
		}
		group := path[1]
		controllerName := path[2]
		if _, ok := c.Group[group]; ok {
			if excontroller, ok2 := c.Group[group][controllerName]; ok2 {
				excontroller.Handlers = append(excontroller.Handlers, h)
			} else {
				c.Group[group][controllerName] = &controller{
					Name:     controllerName,
					Handlers: []*handler{h},
				}
			}
		} else {
			c.Group[group] = map[string]*controller{
				controllerName: {
					Name:     controllerName,
					Handlers: []*handler{h},
				},
			}
		}
		h = nil
	}
	return nil
}
