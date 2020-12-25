package main

import (
	"log"
	"net/http"

	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/logger"
)

func main() {
	s := rest.NewServer("localhost", 8080)
	s.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/user",
		Handler: func(context rest.Context) rest.ErrorInterface {
			logger.ErrorfWithStackWithContext(context, "err : %s", "wow")
			return rest.StatusInternalServerError
		},
	})
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

func hahah() {
	haha2()
}
func haha2() {
	logger.Infof("ok ..... ")
	logger.InfofWithStack("ok ..... ")
}
