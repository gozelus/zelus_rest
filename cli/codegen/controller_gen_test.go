package codegen

import (
	"os"
	"path/filepath"
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
	baseDir := "/Users/zhengli/workspace/private/projects/zelus_rest/cli/codegen/api"
	err = NewControllerGenner(apiReader, baseDir).GenCode()
	if err != nil {
		t.Fatal(err)
	}
	varsFile, err := os.Create(filepath.Join(baseDir, "vars.go"))
	if err != nil {
		t.Fatal(err)
	}
	if err := NewTypesInfo(varsFile, apiReader, "api").GenCode(); err != nil {
		t.Fatal(err)
	}
}
