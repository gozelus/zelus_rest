package main

import (
	"flag"
	"log"

	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/id_generator/internal/routes"
)

var (
	addr     = flag.Int("addr", 8080, "server listened address")
)

func main() {
	s := rest.NewServer(*addr)
	if err := s.AddRoute(routes.Routes...); err != nil {
		log.Fatal(err)
	}
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
