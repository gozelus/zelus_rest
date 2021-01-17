package tpls

var RepoNotFoundErrors = `package {{ .PkgName }}
import (
	"errors"
	"gorm.io/gorm"
)

type RecordErrorNotFound struct {
	Sql string
	Err error
}

func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
`
var RepoNewFuncTpl = `
func New{{.RepoImpName}}(db db.MySQLDb) *{{.RepoImpName}} {
	return &{{.RepoImpName}}{db: db}
}`
var RepoStructTpl = `
type {{.RepoImpName}} struct {
	db db.MySQLDb
}
`
var RepoListFuncTpl = `
// List{{.SelectField.Name}}By{{range .WhereFields}}{{.Name}}{{end}}OrderBy{{.OrderField.Name}}ByTx 根据索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) List{{.SelectField.Name}}By{{range .WhereFields}}{{.Name}}{{end}}OrderBy{{.OrderField.Name}}ByTx(ctx rest.Context,tx db.MySQLDb, {{range .WhereFields}}{{.LowCamelName}} {{.TypeName}},{{end}} limit int64, {{.OrderField.LowCamelName}} {{.OrderField.TypeName}}) ([]*{{.ModelPkgName}}.{{.ModelName}}, bool, error) {
	var resp []*{{.ModelPkgName}}.{{.ModelName}}
	var hasMore bool
	if err := tx.Table(ctx, "{{.TableName}}").
		Select("{{.SelectField.DbName}}, {{ .OrderField.DbName }}").{{range .WhereFields}}
		Where("{{.DbName}} = ?", {{.LowCamelName}}).{{end}}
		Where("{{.OrderField.DbName}} < ?", {{.OrderField.LowCamelName}}).
		Order("{{.OrderField.DbName}} desc").
		Limit(int(limit + 1)).
		Find(&resp); err != nil {
		return nil, false, errors.Wrap(err, "failed in repos")
	}
	hasMore = len(resp) > int(limit)
	if hasMore {
		resp = resp[:len(resp)-1]
	}
	return resp, hasMore, nil
}

// List{{.SelectField.Name}}By{{range .WhereFields}}{{.Name}}{{end}}OrderBy{{.OrderField.Name}} 根据索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) List{{.SelectField.Name}}By{{range .WhereFields}}{{.Name}}{{end}}OrderBy{{.OrderField.Name}}(ctx rest.Context, {{range .WhereFields}}{{.LowCamelName}} {{.TypeName}},{{end}} limit int64, {{.OrderField.LowCamelName}} {{.OrderField.TypeName}}) ([]*{{.ModelPkgName}}.{{.ModelName}}, bool, error) {
	var resp []*{{.ModelPkgName}}.{{.ModelName}}
	var hasMore bool
	if err := repo.db.Table(ctx, "{{.TableName}}").
		Select("{{.SelectField.DbName}}, {{ .OrderField.DbName }}").{{range .WhereFields}}
		Where("{{.DbName}} = ?", {{.LowCamelName}}).{{end}}
		Where("{{.OrderField.DbName}} < ?", {{.OrderField.LowCamelName}}).
		Order("{{.OrderField.DbName}} desc").
		Limit(int(limit + 1)).
		Find(&resp); err != nil {
		return nil, false, errors.Wrap(err, "failed in repos")
	}
	hasMore = len(resp) > int(limit)
	if hasMore {
		resp = resp[:len(resp)-1]
	}
	return resp, hasMore, nil
}
`
var RepoFindManyFuncTpl = `{{$firstField := first .Fields}} {{ $remainFields := (slice .Fields 1) }}
// FindManyWith{{$firstField.Name}}ByTx 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) FindManyWith{{$firstField.Name}}ByTx(ctx rest.Context, tx db.MySQLDb, {{$firstField.LowCamelName}}s []{{$firstField.TypeName}}, {{ range $remainFields }} {{ .LowCamelName }} {{ .TypeName }} {{ end }}) (map[{{$firstField.TypeName}}]*{{.ModelPkgName}}.{{.ModelName}}, error) { 
	resp := map[{{$firstField.TypeName}}]*{{.ModelPkgName}}.{{.ModelName}}{}
	var results []*{{.ModelPkgName}}.{{.ModelName}}
	db := tx.Table(ctx, "{{.TableName}}").
        Where("{{$firstField.DbName}} in (?)", {{$firstField.LowCamelName}}s) 
	{{ range $remainFields }} db = db.Where("{{ .DbName }} = ?", {{ .LowCamelName }}) {{ end }}
	if err := db.Find(&results); err != nil {
		return nil, errors.Wrap(err, "failed in repos")
	}
	for _, r := range results {
		resp[r.{{$firstField.Name}}] = r
	}
	return resp, nil
}
// FindManyWith{{$firstField.Name}} 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) FindManyWith{{$firstField.Name}}(ctx rest.Context, {{$firstField.LowCamelName}}s []{{$firstField.TypeName}}, {{ range $remainFields }} {{ .LowCamelName }} {{ .TypeName }} {{ end }}) (map[{{$firstField.TypeName}}]*{{.ModelPkgName}}.{{.ModelName}}, error) { 
	resp := map[{{$firstField.TypeName}}]*{{.ModelPkgName}}.{{.ModelName}}{}
	var results []*{{.ModelPkgName}}.{{.ModelName}}
	db := repo.db.Table(ctx, "{{.TableName}}").
        Where("{{$firstField.DbName}} in (?)", {{$firstField.LowCamelName}}s)
	{{ range $remainFields }} db = db.Where("{{ .DbName }} = ?", {{ .LowCamelName }}) {{ end }}
	if err := db.Find(&results); err != nil {
		return nil, errors.Wrap(err, "failed in repos")
	}
	for _, r := range results {
		resp[r.{{$firstField.Name}}] = r
	}
	return resp, nil
}
`

var RepoFirstOrCreateFuncTpl = `
// FirstOrCreateWith{{range .Fields}}{{.Name}}{{end}}ByTx 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) FirstOrCreateWith{{range .Fields}}{{.Name}}{{end}}ByTx(ctx rest.Context, tx db.MySQLDb, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}} data *{{.ModelPkgName}}.{{.ModelName}}) error { 
	resp := data {{$lastName := (last .Fields).Name}}
	db := tx.Table(ctx, "{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.Wrap(err, "failed in repos")
	}
	return nil
} 
// FirstOrCreateWith{{range .Fields}}{{.Name}}{{end}} 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) FirstOrCreateWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}} data *{{.ModelPkgName}}.{{.ModelName}}) error { 
	resp := data {{$lastName := (last .Fields).Name}}
	db := repo.db.Table(ctx, "{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.Wrap(err, "failed in repos")
	}
	return nil
} 
`
var RepoFindOneFuncTpl = `
// FindOneWith{{range .Fields}}{{.Name}}{{end}}ByTx 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) FindOneWith{{range .Fields}}{{.Name}}{{end}}ByTx(ctx rest.Context, tx db.MySQLDb, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}}) (*{{.ModelPkgName}}.{{.ModelName}}, error) { 
	resp := &{{.ModelPkgName}}.{{.ModelName}}{} {{$lastName := (last .Fields).Name}}
	db := tx.Table(ctx, "{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.First(resp); err != nil {
		return nil, errors.Wrap(err, "failed in repos")
	}
	return resp, nil
}
// FindOneWith{{range .Fields}}{{.Name}}{{end}}ByTx 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) FindOneWith{{range .Fields}}{{.Name}}{{end}}ByTxForUpdate(ctx rest.Context, tx db.MySQLDb, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}}) (*{{.ModelPkgName}}.{{.ModelName}}, error) { 
	resp := &{{.ModelPkgName}}.{{.ModelName}}{} {{$lastName := (last .Fields).Name}}
	db := tx.Table(ctx, "{{$.TableName}}").Clauses(clause.Locking{Strength: "UPDATE"}).{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.First(resp); err != nil {
		return nil, errors.Wrap(err, "failed in repos")
	}
	return resp, nil
}
// FindOneWith{{range .Fields}}{{.Name}}{{end}} 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) FindOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}}) (*{{.ModelPkgName}}.{{.ModelName}}, error) { 
	resp := &{{.ModelPkgName}}.{{.ModelName}}{} {{$lastName := (last .Fields).Name}}
	db := repo.db.Table(ctx, "{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.First(resp); err != nil {
		return nil, errors.Wrap(err, "failed in repos")
	}
	return resp, nil
}
`
var RepoDeleteFuncTpl = `{{$lastName := (last .Fields).Name}}
// DeleteOneWith{{range .Fields}}{{.Name}}{{end}}ByTx 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) DeleteOneWith{{range .Fields}}{{.Name}}{{end}}ByTx(ctx rest.Context, tx db.MySQLDb, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}}) error { 
	db := tx.Table(ctx, "{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.Delete({{.ModelPkgName}}.{{.ModelName}}{});err!=nil{
		return errors.Wrap(err, "failed in repos")
	}
	return nil
} 
// DeleteOneWith{{range .Fields}}{{.Name}}{{end}} 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) DeleteOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}}) error { 
	db := repo.db.Table(ctx, "{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.Delete({{.ModelPkgName}}.{{.ModelName}}{});err!=nil{
		return errors.Wrap(err, "failed in repos")
	}
	return nil
} 
`
var RepoUpdateFuncTpl = `{{$lastName := (last .Fields).Name}}
// UpdateOneWith{{range .Fields}}{{.Name}}{{end}}ByTx 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) UpdateOneWith{{range .Fields}}{{.Name}}{{end}}ByTx(ctx rest.Context, tx db.MySQLDb, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}} attr map[string]interface{}) error { 
	db := tx.Table(ctx, "{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.Updates(attr);err!=nil{
		return errors.Wrap(err, "failed in repos")
	}
	return nil
} 

// UpdateOneWith{{range .Fields}}{{.Name}}{{end}} 根据唯一索引 {{.IdxName}} 生成
func (repo *{{.RepoImpName}}) UpdateOneWith{{range .Fields}}{{.Name}}{{end}}(ctx rest.Context, {{range .Fields}}{{.LowCamelName}} {{.TypeName}},{{end}} attr map[string]interface{}) error { 
	db := repo.db.Table(ctx, "{{$.TableName}}").{{range $i, $field := .Fields}}
        Where("{{$field.DbName}} = ?", {{$field.LowCamelName}}){{if not (eq $lastName $field.Name)}}.{{end}}{{end}}
	if err := db.Updates(attr);err!=nil{
		return errors.Wrap(err, "failed in repos")
	}
	return nil
} 
`
var RepoCreateFuncTpl = `
// InsertByTx 默认生成的创建函数, 使用 tx 句柄
func (repo *{{.RepoImpName}}) InsertByTx(ctx rest.Context, tx db.MySQLDb, data *{{.ModelPkgName}}.{{.ModelName}}) error {
	if err := tx.Table(ctx, "{{.TableName}}").Insert(data);err!=nil{
		return errors.Wrap(err, "failed in repos")
    }
	return nil
}
// Insert 默认生成的创建函数
func (repo *{{.RepoImpName}}) Insert(ctx rest.Context, data *{{.ModelPkgName}}.{{.ModelName}}) error {
	if err := repo.db.Table(ctx, "{{.TableName}}").Insert(data);err!=nil{
		return errors.Wrap(err, "failed in repos")
    }
	return nil
}
`
