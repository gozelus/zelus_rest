//+build wireinject

package injector

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/gozelus/zelus_rest/example/internal/controller"
	"github.com/gozelus/zelus_rest/example/internal/router"
	"github.com/gozelus/zelus_rest/example/internal/service"
	"github.com/gozelus/zelus_rest/example/repo"
)

var set = wire.NewSet(
	wire.Bind(new(service.UserRepoInter), new(*repo.UserRepo)),
	wire.Bind(new(controller.UserServiceInter), new(*service.UserService)),
	wire.Bind(new(router.UserControllerInter), new(*controller.UserController)),

	repo.NewUserRepo,
	service.NewUserService,
	controller.NewUserController,
	router.NewRouter,
	newdb,
)

func newdb() *sql.DB {
	return nil
}
func NewRouter() *router.Router {
	wire.Build(set)
	return nil
}
