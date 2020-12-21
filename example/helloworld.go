package main

import (
	"fmt"
	"net/http"

	rest "github.com/gozelus/zelus_rest"
)

func main() {
	s := rest.NewServer()
	s.Run("localhost", 8080, struct {
		Method  string
		Path    string
		Handler http.HandlerFunc
	}{Method: http.MethodGet, Path: "/", Handler: func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("hi")
		w.Write([]byte("hi"))
	}})
}
