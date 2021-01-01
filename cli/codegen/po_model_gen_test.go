package codegen

import (
	"os"
	"testing"
)

func TestPoModelStructInfo_GenCode(t *testing.T) {
	m := NewPoModelStructInfo("episode_like_relations", "zhengli:Zhengli_0220@tcp(rm-2ze0o33s1so634285bo.mysql.rds.aliyuncs.com:3306)/podcast?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", "models")
	os.Mkdir("models", 0777)
	os.Mkdir("repos", 0777)
	file, _ := os.Create("./models/models.go")
	if err := m.GenCode(file); err != nil {
		t.Fatalf("%+v", err)
	}
	file2, _ := os.Create("./repos/repo.go")
	if err := NewRepoGener(file2, m, "repos").GenCode(); err != nil {
		t.Fatalf("%+v", err)
	}
}
