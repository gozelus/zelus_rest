package actions

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

func forceCreateFile(path string) (*os.File, error) {
	_, err := os.Stat(path)
	var file *os.File
	if err == nil {
		fmt.Println(color.HiRedString("%s exist, will remove and recreate", path))
		if err = os.Remove(path); err != nil {
			return nil, err
		}
		if file, err = os.Create(path); err != nil {
			return nil, err
		}
	}
	if os.IsNotExist(err) {
		fmt.Println(color.HiGreenString("%s not exist, will recreate", path))
		if file, err = os.Create(path); err != nil {
			return nil, err
		}
	}
	return file, nil
}
func forceCreateDir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println(color.HiRedString("%s exist, will remove and recreate", path))
		if err := os.RemoveAll(path); err != nil {
			return err
		}
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	if os.IsNotExist(err) {
		fmt.Println(color.HiGreenString("%s not exist, will recreate", path))
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func createIfNotExist(path string) (*os.File, bool, error) {
	var file *os.File
	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		fmt.Println(color.HiGreenString("%s is not exist, will create", path))
		if file, err = os.Create(path); err != nil {
			return nil, false, err
		}
		return file, false, nil
	}
	if file, err = os.Open(path); err != nil {
		return nil, true, err
	}
	fmt.Println(color.HiMagentaString("%s exist, will ignore to create", path))
	return file, true, nil
}

func mkdirIfNotExist(dir string) (string, error) {
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(dirAbs); os.IsNotExist(err) {
		fmt.Println(color.HiGreenString("%s is not exist, will create", dirAbs))
		if err := os.MkdirAll(dirAbs, os.ModePerm); err != nil {
			return "", err
		}
	} else {
		fmt.Println(color.HiMagentaString("%s exist, will ignore", dirAbs))
	}
	return dirAbs, nil
}
