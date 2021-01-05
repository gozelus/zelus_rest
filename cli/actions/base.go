package actions

import (
	"fmt"
	"github.com/fatih/color"
	"os/exec"
	"time"
)

func logFinishAndFmt(dirPath string) error {
	now := time.Now()
	fmt.Println(color.HiGreenString("the file : %s will be fmt ...", dirPath))
	if err := exec.Command("goimports", "-w", dirPath).Run(); err != nil {
		return err
	}
	fmt.Println(color.HiGreenString("the file : %s fmt done ---> %dms", dirPath, time.Now().Sub(now).Milliseconds()))
	return nil
}
