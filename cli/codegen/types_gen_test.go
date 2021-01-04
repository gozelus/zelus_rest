package codegen

import (
	"os"
	"testing"
)

func TestTypes(t *testing.T) {
	w, err := os.Create("/Users/zhengli/Desktop/x.go")
	if err != nil {
		t.Fatal(err)
	}
	r, err := os.Open("/Users/zhengli/Desktop/id_generator/api/server.api")
	if err != nil {
		t.Fatal(err)
	}
	if err := NewTypesInfo(w, r, "api").GenCode(); err != nil {
		t.Fatal(err)
	}
}
