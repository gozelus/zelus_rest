package codegen

import (
	"io"
	"text/template"
)

var MainTpl = `package main

import (
	"fmt"
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/logger"
	"{{ .ModuleName }}/internal/routes"
	"{{ .ModuleName }}/config"
	"reflect"
	"runtime"
)

func main() {
	s := rest.NewServer(config.Cfg.Port, func(imp *rest.Plugin) {
		imp.JwtAk = func() (s string, i int64, i2 int64) {
			return config.Cfg.Jwt.Key, config.Cfg.Jwt.Expire, config.Cfg.Jwt.MinTimeToRefresh
		}
	})
	info := ""
	for _, r := range routes.Routes {
		funcName := runtime.FuncForPC(reflect.ValueOf(r.Handler).Pointer()).Name()
		info += fmt.Sprintf("\n[%s] [%s] -> [%s]\n", r.Method, r.Path, funcName)
	}
	fmt.Println(info)

	if err := s.AddRoute(routes.Routes...); err != nil {
		logger.Errorf("add route err for : %s", err)
		os.Exit(1)
	}

	fmt.Printf("will listen in : %d  .... \n", config.Cfg.Port)

	if err := s.Run(); err != nil {
		logger.Errorf("add route err for : %s", err)
		os.Exit(1)
	}
}
`

type MainGenner struct {
	ModuleName string
	AppName    string
}

func NewMainGenner(moduleName string, appName string) *MainGenner {
	return &MainGenner{
		ModuleName: moduleName,
		AppName:    appName,
	}
}

func (m *MainGenner) GenCode(w io.Writer) error {
	var t *template.Template
	var err error
	if t, err = template.New("main gen").Parse(MainTpl); err != nil {
		return err
	}
	return t.Execute(w, m)
}
