package codegen

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

type ApiGenner struct {
	infile    *os.File
	dirpath   string
	collector []*os.File
}

func NewApiGenner(dirpath string, infile *os.File) *ApiGenner {
	return &ApiGenner{
		dirpath:   dirpath,
		infile:    infile,
		collector: []*os.File{infile},
	}
}

func (a *ApiGenner) Merge() (io.Reader, error) {
	if err := a.read(a.infile); err != nil {
		return nil, err
	}
	// merge
	writer := bytes.NewBuffer([]byte(``))
	for i := len(a.collector) - 1; i > 0; i-- {
		file := a.collector[i]
		if _, err := bufio.NewReader(file).WriteTo(writer); err != nil {
			return nil, err
		}
	}
	return writer, nil
}

func (a *ApiGenner) read(file *os.File) error {
	reader := bufio.NewReader(file)
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
			if err := a.readAndMerge(strings.ReplaceAll(strings.Split(lineStr, " ")[1], `"`, "")); err != nil {
				return err
			}
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
