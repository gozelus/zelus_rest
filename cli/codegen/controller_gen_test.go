package codegen

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestGen(t *testing.T) {
	basePath := "/Users/zhengli/workspace/private/projects/zelus_rest/example/id_generator/api"
	apiFile, _ := os.Open(basePath + "/id_generator.api")
	api := NewApiGenner(apiFile)
	apiReader, err := api.Merge()
	if err != nil {
		t.Fatal(err)
	}
	w := bytes.NewBufferString("")
	err = NewControllerGenner(apiReader, w).GenCode()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(w.String())
}
