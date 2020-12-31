package main

var repoTpl = `package models
import (
	"gorm.io/gorm"
	"github.com/pkg/errors"
	"time"
)
type {{.ModelName}}Repo struct {
	db *gorm.DB
}

type {{.ModelName}}Model struct { {{range .Fields}}
    {{.Name}} {{.TypeName}} ` + "`gorm:\"{{.UnderLineName}}\"`" + `{{end}}
}

{{range .IdxKey}}

{{if .IsUniq}}
func (repo *{{$.ModelName}}Repo) FindOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.SmallCamelName}} {{.TypeName}},{{end}}) (*{{$.ModelName}}Model, error) { 
	resp := &{{$.ModelName}}Model{} {{$lastName := (last .Fields).Name}}
	db := repo.db.WithContext(ctx).Table("{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.UnderLineName}} = ?", {{$field.SmallCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
} 
func (repo *{{$.ModelName}}Repo) UpdateOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.SmallCamelName}} {{.TypeName}},{{end}} attr map[string]interface{}) error { 
} 
func (repo *{{$.ModelName}}Repo) DeleteOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.SmallCamelName}} {{.TypeName}},{{end}}) error { 
} 
{{end}}

{{if eq (len .Fields) 1 }}
{{with $field := (index .Fields 0)}}
func (repo *{{$.ModelName}}Repo) FindManyWith{{$field.Name}}s(ctx rest.Context, {{$field.SmallCamelName}}s []{{$field.TypeName}}) (map[{{$field.TypeName}}]*{{$.ModelName}}Model, error) {
} 
func (repo *{{$.ModelName}}Repo) DeleteManyWith{{$field.Name}}s(ctx rest.Context, {{$field.SmallCamelName}}s []{{$field.TypeName}}) error {
} 
{{end}} 
{{end}}
{{end}}
`
