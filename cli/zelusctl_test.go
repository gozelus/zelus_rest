package main

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestGenGoCode(t *testing.T) {
	//GenGoCode()
	GenGoModelCode("zhengli:Zhengli_0220@tcp(rm-2ze0o33s1so634285bo.mysql.rds.aliyuncs.com:3306)/podcast?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
		"episode_like_relations")

	fmt.Printf("%s \n %s \n %+v", Y(), Y(), errors.Cause(Y()))
}
func Y() error {
	return errors.WithMessage(errors.WithStack(errors.New("123")), "hi")
}
