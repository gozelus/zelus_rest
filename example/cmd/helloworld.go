package main

import (
	"log"
	"net/http"
	"time"

	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/logger"
)

func main() {
	s := rest.NewServer("localhost", 8080)
	s.Use(func(context rest.Context) error {
		now := time.Now()
		context.Next()
		logger.Debugf("duration : %v", time.Now().Sub(now))
		return nil
	}, func(context rest.Context) error {
		logger.Debugf("req come ....")
		return nil
	})
	s.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/user",
		Handler: func(context rest.Context) error {
			context.OkJSON("hah")
			return nil
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
