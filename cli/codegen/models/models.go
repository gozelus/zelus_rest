package models
type EpisodeLikeRelationsModel struct { 
 	Id int64 `gorm:"id"`
 	EpisodeId int64 `gorm:"episode_id"`
 	UserId int64 `gorm:"user_id"`
 	CreateTime time.Time `gorm:"create_time"`
 	UpdateTime time.Time `gorm:"update_time"` 
}