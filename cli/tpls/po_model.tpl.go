package tpls

var ModelPackageTpl = `package {{.PackageName}}`
var ModelImportTpl = `import "time"`
var ModelDefineTpl = `
type {{.ModelName}} struct { {{range .Fields}}
 	{{.Name}} {{.TypeName}} ` + "`gorm:\"{{.DbName}}\"`" + `{{end}} 
}`
