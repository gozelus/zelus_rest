package actions

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gozelus/zelus_rest/cli/codegen"
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

	varsFile, err := os.Create(filepath.Join(dir, "vars.go"))
	if err != nil {
		return err
	}
	fmt.Println(color.GreenString("%s create", filepath.Join(dir, "vars.go")))
	if err = codegen.NewTypesInfo(varsFile, apiFileMerge, "api").GenCode(); err != nil {
		return err
	}
	if err := codegen.NewControllerGenner(apiFileMergeCopy, dir, "api").GenCode(); err != nil {
		return err
	}
	return nil
}
