package main

import (
	"github.com/gozelus/zelus_rest"
	_ "gorm.io/driver/mysql"
	"log"
	"net/http"
)

func main() {
	s := rest.NewServer("localhost", 8080)
	if err := s.AddRoute(rest.Route{
		Path:   "/user/create",
		Method: http.MethodGet,
		Handler: func(context rest.Context) {

		},
	}); err != nil {
		log.Fatal(err)
	}
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
