package codegen

import (
	"io"
	"text/template"
)

var CfgTpls = `package config

import (
	"github.com/gozelus/zelus_rest/logger"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Port  int ` + "`" + `yaml:"Port"` + "`" +
	`Mysql struct {
		DataSource string ` + "`" + `yaml:"DataSource"` + "`" + `
	} ` + "`" + `yaml:"Mysql"` + "`" + `
}

var Cfg *Config

func init() {
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
	if err := yaml.NewDecoder(configYaml).Decode(c); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	Cfg = c
}`

type ConfigGenner struct {
	AppName string
}

func NewConfigGenner(appName string) *ConfigGenner {
	return &ConfigGenner{
		AppName: appName,
	}
}

func (c *ConfigGenner) GenCode(w io.Writer) error {
	var t *template.Template
	var err error
	if t, err = template.New("config gen").Parse(CfgTpls); err != nil {
		return err
	}
	return t.Execute(w, c)
}
