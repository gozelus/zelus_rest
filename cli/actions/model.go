package actions

import (
	"github.com/gozelus/zelus_rest/cli/codegen"
	"github.com/urfave/cli"
	"os"
	"strings"
)

const (
	flagDir     = "dir"
	flagUrl     = "url"
	flagTable   = "table"
	flagPkgName = "pkg"
)

func GenModel(ctx *cli.Context) error {
	url := strings.TrimSpace(ctx.String(flagUrl))
	dir := strings.TrimSpace(ctx.String(flagDir))
	pattern := strings.TrimSpace(ctx.String(flagTable))
	pkg := strings.TrimSpace(ctx.String(flagPkgName))

	var dirAbsPath string
	var err error
	var file *os.File

	if dirAbsPath, err = mkdirIfNotExist(dir); err != nil {
		return err
	}
	if file, err = os.Create(dirAbsPath + "/" + pattern + "_model.go"); err != nil {
		return err
	}

	m := codegen.NewPoModelStructInfo(pattern, url, pkg)

	if err = m.GenCode(file); err == nil {
		return logFinishAndFmt(dirAbsPath)
	}
	return err
}
