package actions

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getModuleName() (string, error) {
	modFile, err := os.Open("./go.mod")
	if os.IsNotExist(err) {
		fmt.Println(color.HiRedString("go.mod file is not exist"))
		return "", err
	}
	reader := bufio.NewReader(modFile)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			return "", err
		}
		if strings.Contains(string(lineBytes), "module") {
			keys := strings.Split(string(lineBytes), " ")
			moduleName := keys[1]
			fmt.Println(color.HiGreenString("module name : %s", moduleName))
			return moduleName, nil
		}
	}
}

func logFinishAndFmt(dirPath string) error {
	now := time.Now()
	fmt.Println(color.HiGreenString("the file : %s will be fmt ...", dirPath))
	if err := exec.Command("goimports", "-w", dirPath).Run(); err != nil {
		return err
	}
	fmt.Println(color.HiGreenString("the file : %s fmt done ---> %dms", dirPath, time.Now().Sub(now).Milliseconds()))
	return nil
}
