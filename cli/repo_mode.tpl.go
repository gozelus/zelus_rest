package main

var repoTpl = `package models
import (
	"github.com/gozelus/zelus_rest"
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

func (repo *{{$.ModelName}}Repo) Insert(ctx rest.Context, data *{{.ModelName}}Model) error { 
	if err := repo.db.WithContext(ctx).Table("{{$.TableName}}").Create(data).Error;err!=nil{
		return errors.WithStack(err)
    }
	return nil
}
{{range .IdxKey}}

{{if .IsUniq}}
func (repo *{{$.ModelName}}Repo) FirstOrCreateWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.SmallCamelName}} {{.TypeName}},{{end}} data *{{$.ModelName}}Model) error { 
	resp := data {{$lastName := (last .Fields).Name}}
	db := repo.db.WithContext(ctx).Table("{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.UnderLineName}} = ?", {{$field.SmallCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
} 
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
	db := repo.db.WithContext(ctx).Table("{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.UnderLineName}} = ?", {{$field.SmallCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.Updates(attr).Error;err!=nil{
		return errors.WithStack(err)
	}
	return nil
} 
func (repo *{{$.ModelName}}Repo) DeleteOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.SmallCamelName}} {{.TypeName}},{{end}}) error { 
	db := repo.db.WithContext(ctx).Table("{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.UnderLineName}} = ?", {{$field.SmallCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.Delete({{$.ModelName}}Model{}).Error;err!=nil{
		return errors.WithStack(err)
	}
	return nil
} 
// 非唯一索引
{{else}}
{{ if lt ( len .Fields ) 3}}
// order by create
{{else}}
// select last where first order by mid
{{$selectField := last .Fields}}
{{$queryField := list}}
{{$orderByField := last .Fields}}
{{$fieldsLength := len .Fields}}
{{range $i, $field:= .Fields}}
{{if eq (add -2 $fieldsLength) $i}}
{{$orderByField = $field}}
{{end}}
{{if and (ne (add -1 $fieldsLength) $i) (ne (add -2 $fieldsLength) $i)}}
{{$queryField = append $queryField $field}} 
{{end}}
{{end}}
func (repo *{{$.ModelName}}Repo) Select{{$selectField.Name}}With{{range $queryField}}{{.Name}}{{end}}OrderBy{{$orderByField.Name}}(ctx rest.Context, {{range $queryField}}{{.SmallCamelName}} {{.TypeName}},{{end}} limit int, lastScore {{$orderByField.TypeName}}) ([]*{{$.ModelName}}Model, error){
	var resp []*{{$.ModelName}}Model {{$lastName := (last $queryField).Name}}
	db := repo.db.WithContext(ctx).Table("{{$.TableName}}").{{range $i, $field := $queryField}}
        Where("{{$field.UnderLineName}} = ?", {{$field.SmallCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.Order("{{$orderByField.UnderLineName}} desc").Limit(limit).Find(&resp).Error;err!=nil{
		return nil, errors.WithStack(err)
    }
	return resp, nil
}
{{end}}
{{end}}
{{end}}
`
