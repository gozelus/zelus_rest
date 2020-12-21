package main

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal"
	"github.com/gozelus/zelus_rest/example/internal/controller"
	"github.com/gozelus/zelus_rest/example/internal/service"
	"github.com/gozelus/zelus_rest/example/repo"
)

func main() {
	s := rest.NewServer()
	repo := repo.NewUserRepo(nil)
	userService := service.NewUserService(repo)
	userController := controller.NewUserController(userService)
	router := internal.NewRouter(userController)
	s.AddRoute(router.Routes()...)
	s.Use(rest.LoggerMiddleware)
	s.Run("localhost", 8080)
}
