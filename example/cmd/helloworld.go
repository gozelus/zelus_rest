package main

import (
	"time"

	"github.com/gozelus/zelus_rest/logger"
)

func main() {
	for {
		hahah()
		time.Sleep(5 * time.Second)
	}
}

func hahah() {
	haha2()
}
func haha2() {
	logger.Infof("ok ..... ")
	logger.InfofWithStack("ok ..... ")
}
