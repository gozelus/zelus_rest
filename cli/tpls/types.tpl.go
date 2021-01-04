package tpls

var TypesTpl = `
{{ range .Types }}
{{ .Comment }}
type {{ .TypeName }} struct { {{ range .Fields }}
  {{ .Name }} {{ .TypeName }} {{ .Tags }} {{ .Comment }} {{ end }}
}
{{ end }}
`
