package codegen

import (
	"io"
	"text/template"
)

var EtcYamlTpl = `
AppName: {{ .AppName }}
Port: 8080
JwtKey: {{ .AppName }}-jwt-key # TODO to replace it
`

type EtcGenner struct {
	AppName string
}

func NewEtcGenner(appName string) *EtcGenner {
	return &EtcGenner{
		AppName: appName,
	}
}

func (c *EtcGenner) GenCode(w io.Writer) error {
	var t *template.Template
	var err error
	if t, err = template.New("config gen").Parse(EtcYamlTpl); err != nil {
		return err
	}
	return t.Execute(w, c)
}
