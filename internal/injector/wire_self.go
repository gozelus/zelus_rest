// +build wireinject

package injector

import (
	"github.com/google/wire"
)

var allSet = wire.NewSet(
	zelusCtlSet,
	// 在这里撰写自定义的 provider 等
)
