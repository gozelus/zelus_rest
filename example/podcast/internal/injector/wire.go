//+build wireinject

package injector

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/gozelus/zelus_rest/example/internal/adapter/controllers"
	"github.com/gozelus/zelus_rest/example/internal/adapter/repository"
	"github.com/gozelus/zelus_rest/example/internal/domain/user"
	"github.com/gozelus/zelus_rest/example/internal/router"
)

var set = wire.NewSet(
	wire.Bind(new(router.UserControllerInter), new(*controllers.Controller)),
	wire.Bind(new(user.Repo), new(*repository.UserRepo)),
	wire.Bind(new(controllers.UserDomain), new(*user.Domain)),

	repository.NewUserPepo,
	user.NewDomain,
	controllers.NewController,
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
