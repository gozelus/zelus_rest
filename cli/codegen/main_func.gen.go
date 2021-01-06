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
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"{{ .ModuleName }}/internal/routes"
	"reflect"
	"runtime"
)

type Config struct {
	Port  int ` + "`" + `yaml:"Port"` + "`" + `
	Mysql struct {
		DataSource string ` + "`" + `yaml:"DataSource"` + "`" + `
	}` + "`" + `yaml:"Mysql"` + "`" + `
}

func main() {
	env := os.Getenv("{{ .AppName }}-env")

	configFileName := "config-dev.yaml"
	if env == "production" {
		configFileName = "config.yaml"
	}
	wd, err := os.Getwd()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	configYaml, err := os.Open(filepath.Join(wd, "/etc", configFileName))
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	c := &Config{}
	if err := yaml.NewDecoder(configYaml).Decode(c);err!=nil{
		logger.Error(err)
		os.Exit(1)
	}

	fmt.Printf("\n run with config file : %s \n", configFileName)
	s := rest.NewServer(c.Port)
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

	fmt.Printf("will listen in : %d  .... \n", c.Port)

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
