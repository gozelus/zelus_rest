package main

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/injector"
)

func main() {
	s := rest.NewServer("localhost", 8080)
	s.AddRoute(injector.NewRouter().Routes()...)
	s.Use(rest.LoggerMiddleware)
	s.Run()
}
