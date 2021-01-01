package actions

import (
	"os"
	"path/filepath"
)

func mkdirIfNotExist(dir string) (string, error) {
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(dirAbs); os.IsNotExist(err) {
		if err := os.Mkdir(dirAbs, os.ModePerm); err != nil {
			return "", err
		}
	}
	return dirAbs, nil
}
