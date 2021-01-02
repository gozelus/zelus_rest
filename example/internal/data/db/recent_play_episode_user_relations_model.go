// Code generated by ZelusCtl. DO NOT EDIT.

package models

import "time"

// RecentPlayEpisodeUserRelationsModel 用于描述用户最近播放的单集与用户的关系
type RecentPlayEpisodeUserRelationsModel struct {
	Id         int64     `gorm:"id"`          // 唯一主键id
	UserId     int64     `gorm:"user_id"`     // 用户id
	EpisodeId  int64     `gorm:"episode_id"`  // 单集id
	CreateTime time.Time `gorm:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"update_time"` // 更新时间
}
