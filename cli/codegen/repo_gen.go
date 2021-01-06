package codegen

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"

	"github.com/Masterminds/sprig"
	"github.com/gozelus/zelus_rest/cli/tpls"
	"github.com/iancoleman/strcase"
	"io"
	"text/template"
)

type RepoGener struct {
	file       io.Writer
	model      *PoModelStructInfo
	funcs      []string
	pkgName    string
	ModuleName string
}

func NewRepoGener(file io.Writer, model *PoModelStructInfo, pkgName, moduleName string) *RepoGener {
	return &RepoGener{
		file:       file,
		model:      model,
		pkgName:    pkgName,
		ModuleName: moduleName,
	}
}

func (i *RepoGener) GenErrorsCode(w io.Writer) error {
	t, err := template.New("repo package").Parse(tpls.RepoNotFoundErrors)
	if err != nil {
		return err
	}
	return t.Execute(w, struct {
		PkgName string
	}{
		PkgName: i.pkgName,
	})
}
func (i *RepoGener) GenCode() error {
	tpl := `
package ` + i.pkgName + `

import (
	"github.com/pkg/errors"
	"github.com/gozelus/zelus_rest"
	"gorm.io/gorm"
	"{{ .ModuleName }}/internal/data/po_models"
)`
	t, _ := template.New("repo package").Parse(tpl)
	if err := t.Execute(i.file, i); err != nil {
		return err
	}
	if err := i.genRepoStruct(); err != nil {
		return err
	}
	if err := i.genNewFunc(); err != nil {
		return err
	}
	if err := i.genListFuncs(); err != nil {
		return err
	}
	if err := i.genFindManyFuncs(); err != nil {
		return err
	}
	if err := i.genFindOneFuncs(); err != nil {
		return err
	}
	if err := i.genFirstOrCreate(); err != nil {
		return err
	}
	if err := i.genDeleteFuncs(); err != nil {
		return err
	}
	if err := i.genUpdateFuncs(); err != nil {
		return err
	}
	if err := i.genInsertFunc(); err != nil {
		return err
	}
	return nil
}

func (i *RepoGener) genNewFunc() error {
	t, err := template.New("repo new func define").Parse(tpls.RepoNewFuncTpl)
	if err != nil {
		return err
	}
	return t.Execute(i.file, struct {
		RepoImpName string
	}{
		RepoImpName: strcase.ToCamel(i.model.ModelName + "RepoImp"),
	})
}

// genRepoStruct 生成repo struct
func (i *RepoGener) genRepoStruct() error {
	t, err := template.New("repo imp struct define").Parse(tpls.RepoStructTpl)
	if err != nil {
		return err
	}
	return t.Execute(i.file, struct {
		RepoImpName string
	}{
		RepoImpName: strcase.ToCamel(i.model.ModelName + "RepoImp"),
	})
}

func (i *RepoGener) genListFuncs() error {
	var t *template.Template
	var err error
	if t, err = template.New("list fun gen").Funcs(sprig.HermeticTxtFuncMap()).Parse(tpls.RepoListFuncTpl); err != nil {
		return err
	}
	for _, idx := range i.model.Idx {
		if !idx.IsPrimary && !idx.IsUniq && len(idx.Fields) >= 3 {
			param := struct {
				IdxName      string
				SelectField  *Field
				OrderField   *Field
				WhereFields  []*Field
				RepoImpName  string
				ModelName    string
				TableName    string
				ModelPkgName string
			}{
				IdxName:      idx.Name,
				SelectField:  idx.Fields[len(idx.Fields)-1],
				OrderField:   idx.Fields[len(idx.Fields)-2],
				WhereFields:  idx.Fields[:len(idx.Fields)-2],
				ModelPkgName: "po_models",
				RepoImpName:  strcase.ToCamel(i.model.ModelName + "RepoImp"),
				ModelName:    i.model.ModelName,
				TableName:    i.model.TableName,
			}
			if err := i.genFunc(t, param); err != nil {
				return err
			}
		}
	}
	return nil
}

// findOne 函数
func (i *RepoGener) genFindOneFuncs() error {
	return i.genFuncByUniqIdx(tpls.RepoFindOneFuncTpl)
}

// findMany函数
func (i *RepoGener) genFindManyFuncs() error {
	return i.genFuncByUniqIdx(tpls.RepoFindManyFuncTpl, true)
}

// firstOrCreate 函数
func (i *RepoGener) genFirstOrCreate() error {
	return i.genFuncByUniqIdx(tpls.RepoFirstOrCreateFuncTpl)
}

// delete 函数
func (i *RepoGener) genDeleteFuncs() error {
	return i.genFuncByUniqIdx(tpls.RepoDeleteFuncTpl)
}

// updateOne 函数
func (i *RepoGener) genUpdateFuncs() error {
	return i.genFuncByUniqIdx(tpls.RepoUpdateFuncTpl)
}

// insert 函数生成
func (i *RepoGener) genInsertFunc() error {
	var t *template.Template
	var err error
	if t, err = template.New("insert fun gen").Funcs(sprig.HermeticTxtFuncMap()).Parse(tpls.RepoCreateFuncTpl); err != nil {
		return err
	}
	param := struct {
		RepoImpName  string
		ModelName    string
		TableName    string
		ModelPkgName string
	}{
		ModelPkgName: "po_models",
		RepoImpName:  strcase.ToCamel(i.model.ModelName + "RepoImp"),
		ModelName:    i.model.ModelName,
		TableName:    i.model.TableName,
	}
	return i.genFunc(t, param)
}

// genFunByUniqIdx 根据唯一索引生成
func (i *RepoGener) genFuncByUniqIdx(tpl string, onlyPrimary ...bool) error {
	var genFunc = func(idx *Idx) error {
		var t *template.Template
		var err error
		if t, err = template.New("gen update func").Funcs(sprig.HermeticTxtFuncMap()).Parse(tpl); err != nil {
			return err
		}
		fields := idx.Fields
		return i.genFunc(t, struct {
			IdxName      string
			Fields       []*Field
			RepoImpName  string
			TableName    string
			ModelName    string
			ModelPkgName string
		}{
			IdxName:      idx.Name,
			ModelPkgName: "po_models",
			ModelName:    strcase.ToCamel(i.model.ModelName),
			TableName:    i.model.TableName,
			RepoImpName:  strcase.ToCamel(i.model.ModelName + "RepoImp"),
			Fields:       fields,
		})
	}
	// 找到唯一索引
	for _, idx := range i.model.Idx {
		if len(onlyPrimary) > 0 {
			if !idx.IsPrimary {
				continue
			}
			if err := genFunc(idx); err != nil {
				return err
			}
			return nil
		}
		if idx.IsUniq {
			// 寻找联合索引or主键
			if len(onlyPrimary) > 0 {
			} else {
				if err := genFunc(idx); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// genFunc func生成实现
func (i *RepoGener) genFunc(t *template.Template, param interface{}) error {
	funcStrBuilder := bytes.NewBufferString("")
	if err := t.Execute(funcStrBuilder, param); err != nil {
		return err
	}

	fmt.Println(color.HiGreenString("will gen func : %s", funcStrBuilder.String()))

	i.funcs = append(i.funcs, funcStrBuilder.String())

	if err := t.Execute(i.file, param); err != nil {
		return err
	}
	return nil
}
