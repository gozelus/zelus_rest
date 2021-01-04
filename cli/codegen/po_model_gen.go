package codegen

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
	"strings"

	"github.com/gozelus/zelus_rest/cli/tpls"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Field struct {
	Name          string
	TypeName      string
	LowCamelName  string
	DbName        string
	MySQLTypeName string
	Comment       string
	Tags          string
}

type Idx struct {
	Fields    []*Field
	Name      string
	IsPrimary bool
	IsUniq    bool
}

type PoModelStructInfo struct {
	ddl string

	TableComment string
	PackageName  string
	ModelName    string
	TableName    string
	Fields       []*Field
	FieldsMap    map[string]*Field
	Idx          []*Idx
}

func (m *PoModelStructInfo) GenCode(file io.Writer) error {
	var t *template.Template
	var err error

	// gen package code
	if t, err = template.New("pkg").Parse(tpls.ModelPackageTpl); err != nil {
		return err
	}
	if err := t.Execute(file, m); err != nil {
		return err
	}

	// gen define code
	if t, err = template.New("define").Parse(tpls.ModelDefineTpl); err != nil {
		return err
	}
	if err := t.Execute(file, m); err != nil {
		return err
	}
	return nil
}
func NewPoModelStructInfo(tableName string, datasource string, packageName string) *PoModelStructInfo {
	db, err := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	type Result struct {
		Table string
		DDL   string
	}
	result := Result{}
	r := db.Raw("show create table " + tableName).Row()
	if err := r.Scan(&result.Table, &result.DDL); err != nil {
		panic(err)
	}
	m := &PoModelStructInfo{
		FieldsMap:   map[string]*Field{},
		TableName:   tableName,
		ddl:         result.DDL,
		ModelName:   strcase.ToCamel(tableName + "Model"),
		PackageName: packageName,
	}
	if err = m.initTableComment(); err != nil {
		panic(err)
	}
	if err = m.initFields(); err != nil {
		panic(err)
	}
	if err = m.initIndexs(); err != nil {
		panic(err)
	}
	return m
}

// 初始化索引
func (m *PoModelStructInfo) initIndexs() error {
	buffer := new(bytes.Buffer)
	buffer.WriteString(m.ddl)
	ddlReader := bufio.NewReader(buffer)
	//	PRIMARY KEY (`id`),
	//	UNIQUE KEY `uniq_idx_episode_id_user_id` (`episode_id`,`user_id`),
	//	KEY `idx_user_id_create_time_episode_id` (`user_id`,`create_time`,`episode_id`)
	var lineNum int
	for {
		var err error
		lineStr, _, err := ddlReader.ReadLine()
		lineNum++
		if lineNum == 1 { // 跳过第一行表名定义
			continue
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		keys := strings.Split(strings.TrimLeft(string(lineStr), " "), " ")
		if keys[0] != "PRIMARY" && keys[0] != "UNIQUE" && keys[0] != "KEY" {
			continue
		}

		index := Idx{}
		if keys[0] == "PRIMARY" {
			index.IsUniq = true
			index.IsPrimary = true
			index.Name = keys[0]
		} else if keys[0] == "UNIQUE" {
			index.IsUniq = true
			index.Name = strings.ReplaceAll(keys[2], "`", "")
		} else {
			index.Name = strings.ReplaceAll(keys[1], "`", "")
		}

		lastChar := keys[len(keys)-1]
		for _, keyName := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(lastChar, "(", ""), ")", ""), ",") {
			keyName = strings.ReplaceAll(keyName, "`", "")
			if f, ok := m.FieldsMap[keyName]; ok {
				index.Fields = append(index.Fields, f)
			}
		}
		m.Idx = append(m.Idx, &index)
	}
}

func (m *PoModelStructInfo) initTableComment() error {
	buffer := new(bytes.Buffer)
	buffer.WriteString(m.ddl)
	ddlReader := bufio.NewReader(buffer)
	var lineNum int
	for {
		var err error
		lineStr, _, err := ddlReader.ReadLine()
		lineNum++
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if strings.Contains(string(lineStr), ") ENGINE=") {
			// end line
			keys := strings.Split(string(lineStr), "'")
			if len(keys) >= 3 { // 找到注释
				m.TableComment = keys[len(keys)-2]
			}
			return nil
		}
	}
	return nil
}

// 初始化 fields 字段
func (m *PoModelStructInfo) initFields() error {
	buffer := new(bytes.Buffer)
	buffer.WriteString(m.ddl)
	ddlReader := bufio.NewReader(buffer)
	var lineNum int
	for {
		var err error
		lineStr, _, err := ddlReader.ReadLine()
		lineNum++
		if lineNum == 1 { // 跳过第一行表名定义
			continue
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		keys := strings.Split(strings.TrimLeft(string(lineStr), " "), " ")
		if keys[0] == "PRIMARY" {
			// 之后是索引定义
			return nil
		}
		f := Field{}
		f.DbName = strings.ReplaceAll(keys[0], "`", "")  // 第一项为字段名
		f.MySQLTypeName = strings.Split(keys[1], "(")[0] // 类型名
		f.LowCamelName = strcase.ToLowerCamel(f.DbName)
		f.Name = strcase.ToCamel(f.DbName)

		// 找到备注
		for i, key := range keys {
			if key == "COMMENT" {
				f.Comment = strings.ReplaceAll(strings.ReplaceAll(keys[i+1], "'", ""), ",", "")
			}
		}
		if f.TypeName, err = getGolangTypeWithMysqlType(f.MySQLTypeName); err != nil {
			return err
		}
		m.FieldsMap[f.DbName] = &f
		m.Fields = append(m.Fields, &f)
	}
}

func getGolangTypeWithMysqlType(mysqlType string) (string, error) {
	switch mysqlType {
	case "tinyint":
		return "int64", nil
	case "int":
		return "int64", nil
	case "smallint":
		return "int8", nil
	case "mediumint":
		return "int64", nil
	case "bigint":
		return "int64", nil
	case "decimal":
		return "float", nil
	case "float":
		return "float", nil
	case "double":
		return "float", nil
	case "datetime":
		return "time.Time", nil
	case "time":
		return "time.Time", nil
	case "timestamp":
		return "time.Time", nil
	case "varchar":
		return "string", nil
	case "longtext":
		return "string", nil
	case "mediumtext":
		return "string", nil
	case "text":
		return "string", nil
	case "tinytext":
		return "string", nil
	}
	return "", errors.WithStack(errors.New("not found with " + mysqlType))
}
