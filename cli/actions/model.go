
package actions

import (
	"github.com/gozelus/zelus_rest/cli/codegen"
	"github.com/urfave/cli"
	"os"
	"strings"
)

const (
	flagUrl   = "url"
	flagApp   = "appname"
	flagTable = "table"
	flagFile  = "file"
)

func GenModel(ctx *cli.Context) error {
	url := strings.TrimSpace(ctx.String(flagUrl))
	pattern := strings.TrimSpace(ctx.String(flagTable))

	var err error
	var file *os.File

	if _, err := mkdirIfNotExist("./internal/data/po_models"); err != nil {
		return err
	}
	for _, table := range strings.Split(pattern, ",") {
		if file, err = forceCreateFile("./internal/data/po_models/" + table + "_po_model.go"); err != nil {
			return err
		}
		m := codegen.NewPoModelStructInfo(table, url, "po_models")
		if err = m.GenCode(file); err != nil {
			return err
		}
		if err := logFinishAndFmt(file.Name()); err != nil {
			return err
		}
	}
	return nil
}
