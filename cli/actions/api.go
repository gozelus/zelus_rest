package actions

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gozelus/zelus_rest/cli/codegen"
	"github.com/iancoleman/strcase"
	"github.com/urfave/cli"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func mergeApiFile(apiFilePath string) (io.Reader, error) {
	var err error
	var apiFileMerge io.Reader
	apiFile, err := os.Open(apiFilePath)
	if err != nil {
		return nil, err
	}
	if apiFileMerge, err = codegen.NewApiGenner(apiFile).Merge(); err != nil {
		return nil, err
	}
	return apiFileMerge, nil
}

func GenApis(ctx *cli.Context) error {
	dir := strings.TrimSpace(ctx.String(flagDir))
	apiFilePath := strings.TrimSpace(ctx.String(flagFile))

	// apiFileMerge 需要重新读取一次
	var apiFileMergeCopy io.Reader
	var apiFileMerge io.Reader
	var err error

	if apiFileMergeCopy, err = mergeApiFile(apiFilePath); err != nil {
		return err
	}
	if apiFileMerge, err = mergeApiFile(apiFilePath); err != nil {
		return err
	}

	// gen api.vars.go
	// 生成api包的类型定义文件
	varsFile, err := os.Create(filepath.Join(dir, "vars.go"))
	if err != nil {
		return err
	}
	fmt.Println(color.GreenString("%s create", filepath.Join(dir, "vars.go")))
	if err = codegen.NewTypesInfo(varsFile, apiFileMerge, "api").GenCode(); err != nil {
		return err
	}
	if err := logFinishAndFmt(filepath.Join(dir, "vars.go")); err != nil {
		return err
	}

	// 先解析 apiFile 以此生成 controllers
	controllerGen := codegen.NewControllerGenner()
	groupControllers, err := controllerGen.ParseApiFile(apiFileMergeCopy, "api")
	if err != nil {
		return err
	}
	// 生成 service 代码，用于服务于 controllers
	for groupName, controllersMap := range groupControllers {
		path := filepath.Join(dir, "services", groupName)
		// 因为 service 层可能会有些业务代码，所以这个地方不再强制生成
		if _, err := mkdirIfNotExist(path); err != nil {
			return err
		}
		// 遍历 controller 准备生成对应的 service 文件
		for _, c := range controllersMap {
			filename := filepath.Join(path, strcase.ToSnake(c.Name+"_service.go"))
			createFile, err := createIfNotExist(filename)
			if err != nil {
				return err
			}
			if createFile == nil {
				fmt.Println(color.HiRedString("%s exist, will ignore ...", filename))
				continue
			}

			fmt.Println(color.HiGreenString("%s created", filename))
			// 交给 genner
			if err := codegen.NewServiceGener(c).GenCode(createFile); err != nil {
				return err
			}
			if err := logFinishAndFmt(createFile.Name()); err != nil {
				return err
			}
		}
	}

	// 生成 controllers 的代码，用于服务 routes
	for groupName, controllersMap := range groupControllers {
		// 查看是否存在 dir/controllers/$groupName 这个文件夹
		// 如果存在，则强制删除，然后创建新的文件夹
		path := filepath.Join(dir, "controllers", groupName)
		if err := forceCreateDir(path); err != nil {
			return err
		}

		// 遍历 controller 准备生成对应的 controller 文件
		for _, c := range controllersMap {
			filename := filepath.Join(path, strcase.ToSnake(c.Name+"_controller.go"))
			w, err := os.Create(filename)
			if err != nil {
				return err
			}
			fmt.Println(color.HiGreenString("%s created", filename))
			// 交给 genner
			if err := controllerGen.GenCode(w, c); err != nil {
				return err
			}
			if err := logFinishAndFmt(w.Name()); err != nil {
				return err
			}
		}
	}
	return nil
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
