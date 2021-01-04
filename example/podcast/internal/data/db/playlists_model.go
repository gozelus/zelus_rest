// Code generated by ZelusCtl. DO NOT EDIT.

package models

import "time"

// PlaylistsModel 播单详情
type PlaylistsModel struct {
	Id          int64     `gorm:"id"`          // 播单唯一id
	CreateTime  time.Time `gorm:"create_time"` // 创建时间
	UpdateTime  time.Time `gorm:"update_time"` // 更新时间
	UserId      int64     `gorm:"user_id"`     // 用户id
	Title       string    `gorm:"title"`       // 播单标题
	Description string    `gorm:"description"` // 播单描述
}