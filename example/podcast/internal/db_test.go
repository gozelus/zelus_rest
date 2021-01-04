package internal

import (
	db2 "github.com/gozelus/zelus_rest/core/db"
	"sync"
	"testing"
	"time"
)

type UsersModel struct {
	Id         int64     `gorm:"id"`          // 用户唯一id
	CreateTime time.Time `gorm:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"update_time"` // 更新时间
	Nickname   string    `gorm:"nickname"`    // 用户昵称
	Avatar     string    `gorm:"avatar"`      // 头像guid
	Sign       string    `gorm:"sign"`        // 用户签名
}

func TestGormStatement(t *testing.T) {
	db, err := db2.Open("zhengli:Zhengli_0220@tcp(rm-2ze0o33s1so634285bo.mysql.rds.aliyuncs.com:3306)/podcast?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for i := 10; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			u := UsersModel{}
			if err := db.Table(nil, "users").Where("id = ?", i).Find(&u); err != nil {
				t.Fatal(err)
			}
			t.Log(u.Nickname)
		}(i)
	}
	wg.Wait()
}
