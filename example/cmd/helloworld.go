package main

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/injector"
)

func main() {
	s := rest.NewServer()
	s.AddRoute(injector.NewRouter().Routes()...)
	s.Use(rest.LoggerMiddleware)
	s.Run("localhost", 8080)
}
