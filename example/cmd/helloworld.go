package main

import (
	"log"
	"net/http"

	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/controller/user"
	"github.com/gozelus/zelus_rest/logger"
)

func main() {
	s := rest.NewServer("localhost", 8080)
	s.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/user",
		Handler: func(context rest.Context) rest.ErrorInterface {
			req := &user.RegisterRequest{}
			context.JSONQueryBind(req)
			logger.InfofWithContext(context, "nickname : %s", req.Nickname)
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
