package codegen

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ApiGenner struct {
	infile    *os.File
	endReader *bytes.Buffer
}

func NewApiGenner(infile *os.File) *ApiGenner {
	return &ApiGenner{
		infile:    infile,
		endReader: bytes.NewBufferString(``),
	}
}

func (a *ApiGenner) Merge() (io.Reader, error) {
	if err := a.read(a.infile); err != nil {
		return nil, err
	}
	// merge
	return a.endReader, nil
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
		} else {
			a.endReader.WriteString(lineStr + "\n")
		}
	}
}

func (a *ApiGenner) readAndMerge(path string) error {
	dirPath := filepath.Dir(a.infile.Name())
	needMerge, err := os.Open(filepath.Join(dirPath, path))
	if err != nil {
		return err
	}
	return a.read(needMerge)
}
