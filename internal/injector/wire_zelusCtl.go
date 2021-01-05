// +build wireinject

package injector

import (
	"github.com/google/wire"

	v1_controllers "github.com/gozelus/zelus_rest/internal/controllers/v1"

	v1_services "github.com/gozelus/zelus_rest/internal/services/v1"
)

// 此 set 为代码生成，请勿改动
// 如果有其他依赖注入需求，应该新建文件而不要改动此文件
var zelusCtlSet = wire.NewSet(
	v1_controllers.NewUserController,
	v1_services.NewUserService,
	v1_controllers.NewEpisodeController,
	v1_services.NewEpisodeService,
)

func NewV1ControllersUserController() *v1_controllers.UserController {
	wire.Build(allSet)
	return nil
}

func NewV1ServicesUserService() *v1_services.UserService {
	wire.Build(allSet)
	return nil
}

func NewV1ControllersEpisodeController() *v1_controllers.EpisodeController {
	wire.Build(allSet)
	return nil
}

func NewV1ServicesEpisodeService() *v1_services.EpisodeService {
	wire.Build(allSet)
	return nil
}
