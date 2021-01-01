package actions

import (
	"fmt"
	"github.com/fatih/color"
	"os/exec"
)

func logFinishAndFmt(dirPath string) error {
	if err := exec.Command("goimports", "-w", dirPath).Run(); err != nil {
		return err
	}
	fmt.Println(color.HiGreenString("the file : %s done \n", dirPath))
	return nil
}
