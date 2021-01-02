package example

import (
	db2 "github.com/gozelus/zelus_rest/core/db"
	"github.com/gozelus/zelus_rest/example/internal/data/db"
	"testing"
)

func TestGormStatement(t *testing.T) {
	db, err := db2.Open("zhengli:Zhengli_0220@tcp(rm-2ze0o33s1so634285bo.mysql.rds.aliyuncs.com:3306)/podcast?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		t.Fatal(err)
	}
	u := models.UsersModel{}
	if err := db.Table(nil, "users").Where("id = ?", 17).Find(&u); err != nil {
		t.Fatal(err)
	}
	t.Log(u.Nickname)
}
