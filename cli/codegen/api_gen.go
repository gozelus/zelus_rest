package codegen

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type ApiGenner struct {
	infile    *os.File
	dirpath   string
	collector []*os.File
	needMerge []*os.File
}

func NewApiGenner(dirpath string, infile *os.File) *ApiGenner {
	return &ApiGenner{
		dirpath:   dirpath,
		infile:    infile,
		collector: []*os.File{infile},
		needMerge: []*os.File{},
	}
}

func (a *ApiGenner) read(file *os.File) error {
	reader := bufio.NewReader(a.infile)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		lineStr := string(lineBytes)
		if strings.HasPrefix(lineStr, "import") {
		}
	}
}

func (a *ApiGenner) readAndMerge(path string) error {
	needMerge, err := os.Open(a.dirpath + path)
	if err != nil {
		return err
	}
	a.collector = append(a.collector, needMerge)
	return nil
}
