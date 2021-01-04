package codegen

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T)  {
	//apiFile, _ := os.Open("/Users/momo/Desktop/zelus_rest/example/id_generator/api/id_generator.api")
	//vars, _ := os.Create("vars.go")
	//info := NewTypesInfo(vars, apiFile, "codegen")
	//if err := info.GenCode();err!=nil{
	//	t.Fatal(err)
	//}
	var s []int = []int{1}
	for {
		for _, i := range s {
			fmt.Println(i)
			s = append(s, i)
		}
	}
}

