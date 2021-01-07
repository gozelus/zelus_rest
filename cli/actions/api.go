package actions

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gozelus/zelus_rest/cli/codegen"
	"github.com/iancoleman/strcase"
	"github.com/urfave/cli"
	"io"
	"os"
	"os/exec"
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
	dir, _ := os.Getwd()
	dir = filepath.Join(dir, "internal")
	if _, err := mkdirIfNotExist(dir); err != nil {
		return err
	}

	apiFilePath := strings.TrimSpace(ctx.String(flagFile))
	appName := strings.TrimSpace(ctx.String(flagApp))
	// 尝试生成 config 文件
	if _, err := mkdirIfNotExist("./config"); err != nil {
		return err
	}
	cfgFile, ex, err := createIfNotExist("./config/cfg.go")
	if err != nil {
		return err
	}
	if !ex {
		if err := codegen.NewConfigGenner(appName).GenCode(cfgFile); err != nil {
			return err
		}
	}

	// apiFileMerge 需要重新读取一次
	var apiFileMergeCopy io.Reader
	var apiFileMerge io.Reader
	var moduleName string

	moduleName, err = getModuleName()
	if err != nil {
		return err
	}

	if apiFileMergeCopy, err = mergeApiFile(apiFilePath); err != nil {
		return err
	}
	if apiFileMerge, err = mergeApiFile(apiFilePath); err != nil {
		return err
	}

	// 生成可执行文件
	if _, err = mkdirIfNotExist("./cmd"); err != nil {
		return err
	}
	mainFile, err := forceCreateFile(filepath.Join("./cmd", appName+".go"))
	if err != nil {
		return err
	}
	if err := codegen.NewMainGenner(moduleName, appName).GenCode(mainFile); err != nil {
		return err
	}
	if err := logFinishAndFmt(mainFile.Name()); err != nil {
		return err
	}

	// docker file
	dockerFile, ex, err := createIfNotExist("./Dockerfile")
	if err != nil {
		return err
	}
	if !ex {
		if err := codegen.NewDockerFileGenner(appName).GenCode(dockerFile); err != nil {
			return err
		}
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
	groupControllers, err := controllerGen.ParseApiFile(apiFileMergeCopy, "api", moduleName)
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
			createFile, exist, err := createIfNotExist(filename)
			if err != nil {
				return err
			}
			if exist {
				fmt.Println(color.HiRedString("%s exist, will ignore to write %s ...", filename, c.Name))
				continue
			}

			fmt.Println(color.HiGreenString("%s created", filename))
			// 交给 genner
			if err := codegen.NewServiceGener(moduleName).GenCode(createFile, c); err != nil {
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

	// 最后生成对应的 wire set
	var controllers []codegen.Controller
	for _, cmap := range groupControllers {
		for _, c := range cmap {
			controllers = append(controllers, *c)
		}
	}
	_, err = mkdirIfNotExist(filepath.Join(dir, "injector"))
	if err != nil {
		return err
	}
	zelusWireFile, err := forceCreateFile(filepath.Join(dir, "injector", "wire_zelusCtl.go"))
	if err != nil {
		return err
	}
	selfWireFile, ex, err := createIfNotExist(filepath.Join(dir, "injector", "wire_self.go"))
	if ex {
		selfWireFile = nil
	}

	if err := codegen.NewWireGenner(controllers, moduleName).GenCode(zelusWireFile, selfWireFile); err != nil {
		return err
	}
	if !ex {
		if err := logFinishAndFmt(selfWireFile.Name()); err != nil {
			return err
		}
	}
	if err := logFinishAndFmt(zelusWireFile.Name()); err != nil {
		return err
	}

	err = exec.Command("wire", "./internal/injector/wire_zelusCtl.go", "./internal/injector/wire_self.go").Run()
	if err != nil {
		fmt.Println(color.HiRedString("wire err for %s", err))
		return err
	}

	if err := forceCreateDir(filepath.Join(dir, "routes")); err != nil {
		return err
	}
	var routesFile *os.File
	if routesFile, err = forceCreateFile(filepath.Join(dir, "routes", "routes.go")); err != nil {
		return err
	}
	if err := codegen.NewRouteGenner(moduleName).GenCode(routesFile, controllers); err != nil {
		return err
	}
	if err := logFinishAndFmt(routesFile.Name()); err != nil {
		return err
	}
	return nil
}
