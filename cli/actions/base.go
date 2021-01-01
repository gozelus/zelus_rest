package actions

import (
	"github.com/gozelus/zelus_rest/logger"
	"os/exec"
)

func logFinishAndFmt(dirPath string) error {
	if err := exec.Command("goimports", "-w", dirPath).Run(); err != nil {
		return err
	}
	logger.Infof("%s done", dirPath)
	return nil
}
