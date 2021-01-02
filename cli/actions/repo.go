package actions

import (
	"github.com/gozelus/zelus_rest/cli/codegen"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func GenRepo(ctx *cli.Context) error {
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

	for _, table := range strings.Split(pattern, ",") {
		if file, err = os.Create(dirAbsPath + "/" + table + "_repo.go"); err != nil {
			return err
		}
		m := codegen.NewPoModelStructInfo(table, url, pkg)
		r := codegen.NewRepoGener(file, m, pkg)
		if err = r.GenCode(); err != nil {
			return err
		}
		if err := logFinishAndFmt(file.Name());err!=nil{
			return err
		}
	}
	return nil
}
