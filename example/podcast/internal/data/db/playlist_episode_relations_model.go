// Code generated by ZelusCtl. DO NOT EDIT.

package models

import "time"

// PlaylistEpisodeRelationsModel 描述播单和单集直接的关系
type PlaylistEpisodeRelationsModel struct {
	Id         int64     `gorm:"id"`          // 主键ID
	CreateTime time.Time `gorm:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"update_time"` // 更新时间
	UserId     int64     `gorm:"user_id"`     // 用户id
	EpisodeId  int64     `gorm:"episode_id"`  // 单集id
	PlaylistId int64     `gorm:"playlist_id"` // 播单id
}