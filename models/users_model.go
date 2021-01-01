package models

import "time"

type UsersModel struct {
	Id         int64     `gorm:"id"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
	UserId     int64     `gorm:"user_id"`
	Nickname   string    `gorm:"nickname"`
	Avatar     string    `gorm:"avatar"`
	Sign       string    `gorm:"sign"`
}
