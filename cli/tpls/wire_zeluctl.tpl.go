package tpls

var WireSelfTpls = `// +build wireinject

package injector

import (
	"github.com/google/wire"
)


var allSet = wire.NewSet(
	zelusCtlSet,
	// 在这里撰写自定义的 provider 等
)
`

var WireZelusTpls = `// +build wireinject

package injector

import (
	"github.com/google/wire"
	{{ range $import, $key := .Imports }}
	{{ $import }}
	{{ end }}
)

// 此 set 为代码生成，请勿改动
// 如果有其他依赖注入需求，应该新建文件而不要改动此文件
var zelusCtlSet = wire.NewSet({{ range .Constructors }}
	{{ .PkgName }}.New{{ .ReturnName }},{{end}}
)
{{ range .Constructors }}
func New{{ .PkgNameCamel }}{{ .ReturnName }}() *{{ .PkgName }}.{{ .ReturnName }} {
	wire.Build(allSet)
	return nil
}
{{ end }}
`