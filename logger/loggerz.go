package logger

import "fmt"

func Infof(str string, args ...interface{}) {
	fmt.Printf(str+"\n", args...)
}
func Debugf(str string, args ...interface{}) {
	fmt.Printf(str+"\n", args...)
}
func Errorf(str string, args ...interface{}) {
	fmt.Printf(str+"\n", args...)
}
