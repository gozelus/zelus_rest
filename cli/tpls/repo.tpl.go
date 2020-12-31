package tpls

var RepoInterfaceTpl = ``
var RepoStructTpl = `
type {{.RepoImpName}} struct {
	db *gorm.DB
}
`
var RepoFindManyFuncTpl = `{{$firstField := first .Fields}}
func (repo *{{.RepoImpName}}) FindManyWith{{$firstField.Name}}(ctx rest.Context, {{$firstField.LowCamelName}}s []{{$firstField.TypeName}}) (map[{{$firstField.TypeName}}]*{{.ModelPkgName}}.{{.ModelName}}, error) { 
	resp := map[{{$firstField.TypeName}}]*{{.ModelPkgName}}.{{.ModelName}}{}
	var results []*{{.ModelPkgName}}.{{.ModelName}}
	db := repo.db.WithContext(ctx).Table("{{.TableName}}").
        Where("{{$firstField.DbName}} in (?)", {{$firstField.LowCamelName}}s)
	if err := db.Find(&results).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	for _, r := range results {
		resp[r.{{$firstField.Name}}] = r
	}
	return resp, nil
}`

var RepoFirstOrCreateFuncTpl = `
func (repo *{{.RepoImpName}}) FirstOrCreateWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}} data *{{.ModelPkgName}}.{{.ModelName}}) error { 
	resp := data {{$lastName := (last .Fields).Name}}
	db := repo.db.WithContext(ctx).Table("{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
} 
`
var RepoFindOneFuncTpl = `
func (repo *{{.RepoImpName}}) FindOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}}) (*{{.ModelPkgName}}.{{.ModelName}}, error) { 
	resp := &{{.ModelPkgName}}.{{.ModelName}}{} {{$lastName := (last .Fields).Name}}
	db := repo.db.WithContext(ctx).Table("{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}
`
var RepoDeleteFuncTpl = `{{$lastName := (last .Fields).Name}}
func (repo *{{.RepoImpName}}) DeleteOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}}) error { 
	db := repo.db.WithContext(ctx).Table("{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.Delete({{.ModelPkgName}}.{{.ModelName}}{}).Error;err!=nil{
		return errors.WithStack(err)
	}
	return nil
} 
`
var RepoUpdateFuncTpl = `{{$lastName := (last .Fields).Name}}
func (repo *{{.RepoImpName}}) UpdateOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}} attr map[string]interface{}) error { 
	db := repo.db.WithContext(ctx).Table("{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.Updates(attr).Error;err!=nil{
		return errors.WithStack(err)
	}
	return nil
} 
`
var RepoInsertFuncTpl = `
func (repo *{{.RepoImpName}}) Insert(ctx rest.Context, data *{{.ModelPkgName}}.{{.ModelName}}) error {
	if err := repo.db.WithContext(ctx).Table("{{.TableName}}").Create(data).Error;err!=nil{
		return errors.WithStack(err)
    }
	return nil
}
`
