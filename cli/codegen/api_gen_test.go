package codegen

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestA(t *testing.T) {
	file, err := os.Open("/Users/momo/Desktop/zelus_rest/example/id_generator/api/id_generator_api.api")

	r, err := NewApiGenner("/Users/momo/Desktop/zelus_rest/example/id_generator/api", file).Merge()
	if err != nil {
		t.Fatal(err)
	}
	r2 := bufio.NewReader(r)
	for {
		line, _, err := r2.ReadLine()
		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}
			return
		}
		fmt.Println(string(line))
	}
}
