package plugins

//JsReqHandler
// js 代码专用请求定义
type JsReqHandler struct {
	Name   string
	Method string
	URL    string
	Params []string
}

type JsCodeGenerator struct {
}

var jsCodeTemplate = `
const axios = require("axios")

{{ range .Handlers }}
async function {{ .Name }}({{ range .Params }} {{ . }} {{ end }}) {
	{{ if eq .Method "GET" }}
	axios.get("{{ .URL }}", 
		params: {
			{{ range .Params }} {{ . }} : {{ . }} {{ end }}
		}	
	)
	{{ end }}
} {{ end }}
`
