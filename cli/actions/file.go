package actions

import (
	"os"
	"path/filepath"
)

func createIfNotExist(path string) (*os.File, error) {
	var file *os.File
	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if file, err = os.Create(path); err != nil {
			return nil, err
		}
		return file, nil
	}
	return nil, nil
}

func mkdirIfNotExist(dir string) (string, error) {
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(dirAbs); os.IsNotExist(err) {
		if err := os.MkdirAll(dirAbs, os.ModePerm); err != nil {
			return "", err
		}
	}
	return dirAbs, nil
}
