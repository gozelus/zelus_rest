package main

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/controller/user"
	"log"
	"net/http"
)

func main() {
	s := rest.NewServer("localhost", 8080)
	uc := user.NewController(nil)
	if err := s.AddRoute(rest.Route{
		Path:    "/user/create",
		Method:  http.MethodGet,
		Handler: uc.Register,
	}); err != nil {
		log.Fatal(err)
	}
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
