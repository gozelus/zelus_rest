package main

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	"github.com/gozelus/zelus_rest/logger"
	_ "gorm.io/driver/mysql"
	"log"
	"net/http"
)

type User struct {
	ID       int64  `gorm:"id"`
	Nickname string `json:"nickname" gorm:"nickname"`
}

func (u *User) Data() interface{} {
	return u
}

func (u *User) ErrorCode() int {
	return 200
}

func (User) ErrorMessage() string {
	return "ok"
}

func main() {
	s := rest.NewServer("localhost", 8080)
	if err := s.AddRoute(rest.Route{
		Path:   "/user/create",
		Method: http.MethodPost,
		Handler: func(context rest.Context) {
			d, _ := db.Open("zhengli:Zhengli_0220@tcp(rm-2ze0o33s1so634285bo.mysql.rds.aliyuncs.com:3306)/podcast?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
			u := User{}
			var _ rest.Rsp = &User{}
			if err := context.JSONBodyBind(&u); err != nil {
				logger.Errorf("err : %s", err)
				context.RenderJSON(nil)
			}
			d.Table(context, "users").Insert(&u)
			context.RenderJSON(&u)
		},
	}, rest.Route{
		Path:   "/user/find",
		Method: http.MethodGet,
		Handler: func(context rest.Context) {
			d, _ := db.Open("zhengli:Zhengli_0220@tcp(rm-2ze0o33s1so634285bo.mysql.rds.aliyuncs.com:3306)/podcast?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
			u := User{}
			var _ rest.Rsp = &User{}
			context.JSONQueryBind(&u)
			d.Table(context, "users").Where("id = ?", u.ID).Find(&u)
			context.RenderJSON(&u)
		},
	}); err != nil {
		log.Fatal(err)
	}
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
