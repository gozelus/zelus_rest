package actions

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gozelus/zelus_rest/cli/codegen"
	"github.com/urfave/cli"
	"strings"
)

func GenRepo(ctx *cli.Context) error {
	url := strings.TrimSpace(ctx.String(flagUrl))
	pattern := strings.TrimSpace(ctx.String(flagTable))
	var err error
	moduleName, err := getModuleName()
	if err != nil {
		return err
	}

	if _, err = mkdirIfNotExist("./internal/biz/repos"); err != nil {
		return err
	}
	errorsFile, ex, err := createIfNotExist("./internal/biz/repos/errors.go")
	if err != nil {
		return err
	}
	if !ex {
		r := codegen.NewRepoGener(nil, nil, "repos", moduleName)
		if err := r.GenErrorsCode(errorsFile); err != nil {
			return err
		}
	}

	for _, table := range strings.Split(pattern, ",") {
		path := fmt.Sprintf("./internal/biz/repos/%s_repo.go", table)
		file, ex, err := createIfNotExist(path)
		if err != nil {
			return nil
		}
		if ex {
			fmt.Println(color.MagentaString("%s repo file exist , will ignore to write ... ", path))
			continue
		}
		m := codegen.NewPoModelStructInfo(table, url, "repos")
		r := codegen.NewRepoGener(file, m, "repos", moduleName)
		if err = r.GenCode(); err != nil {
			return err
		}
		if err := logFinishAndFmt(file.Name()); err != nil {
			return err
		}
	}
	return nil
}
