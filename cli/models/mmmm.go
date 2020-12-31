package models
import (
	"github.com/gozelus/zelus_rest"
	"gorm.io/gorm"
	"github.com/pkg/errors"
	"time"
)
type EpisodeLikeRelationsRepo struct {
	db *gorm.DB
}

type EpisodeLikeRelationsModel struct { 
    Id int64 `gorm:"id"`
    EpisodeId int64 `gorm:"episode_id"`
    UserId int64 `gorm:"user_id"`
    CreateTime time.Time `gorm:"create_time"`
    UpdateTime time.Time `gorm:"update_time"`
}




func (repo *EpisodeLikeRelationsRepo) FindOneWithId(ctx rest.Context, id int64,) (*EpisodeLikeRelationsModel, error) { 
	resp := &EpisodeLikeRelationsModel{} 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("id = ?", id)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
} 
func (repo *EpisodeLikeRelationsRepo) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error { 
} 
func (repo *EpisodeLikeRelationsRepo) DeleteOneWithId(ctx rest.Context, id int64,) error { 
} 




func (repo *EpisodeLikeRelationsRepo) FindManyWithIds(ctx rest.Context, ids []int64) (map[int64]*EpisodeLikeRelationsModel, error) {
} 
func (repo *EpisodeLikeRelationsRepo) DeleteManyWithIds(ctx rest.Context, ids []int64) error {
} 
 




func (repo *EpisodeLikeRelationsRepo) FindOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64,userId int64,) (*EpisodeLikeRelationsModel, error) { 
	resp := &EpisodeLikeRelationsModel{} 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("episode_id = ?", episodeId).
        Where("user_id = ?", userId)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
} 
func (repo *EpisodeLikeRelationsRepo) UpdateOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64,userId int64, attr map[string]interface{}) error { 
} 
func (repo *EpisodeLikeRelationsRepo) DeleteOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64,userId int64,) error { 
} 









