
package repos

import (
	"github.com/gozelus/zelus_rest/cli/codegen/models"
	"github.com/pkg/errors"
	"github.com/gozelus/zelus_rest"
	"gorm.io/gorm"
	"time"
)
type EpisodeLikeRelationsModelRepoImp struct {
	db *gorm.DB
}

// ListEpisodeIdByUserIdOrderByCreateTime 根据索引 idx_user_id_create_time_episode_id 生成
func (repo *EpisodeLikeRelationsModelRepoImp) ListEpisodeIdByUserIdOrderByCreateTime(ctx rest.Context, userId int64, limit int64, createTime time.Time) ([]*models.EpisodeLikeRelationsModel, bool, error) {
	var resp []*models.EpisodeLikeRelationsModel
	var hasMore bool
	if err := repo.db.WithContext(ctx).Table("episode_like_relations").
		Select("episode_id").
		Where("user_id = ?", userId).
		Where("create_time < ?", createTime).
		Order("create_time desc").
		Limit(int(limit + 1)).
		Find(&resp).Error; err != nil {
		return nil, false, errors.WithStack(err)
	}
	hasMore = len(resp) > int(limit)
	if hasMore {
		resp = resp[:len(resp)-1]
	}
	return resp, hasMore, nil
}

// FindManyWithId(ctx rest.Context, ids []int64 根据唯一索引 PRIMARY 生成
func (repo *EpisodeLikeRelationsModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.EpisodeLikeRelationsModel, error) { 
	resp := map[int64]*models.EpisodeLikeRelationsModel{}
	var results []*models.EpisodeLikeRelationsModel
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("id in (?)", ids)
	if err := db.Find(&results).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	for _, r := range results {
		resp[r.Id] = r
	}
	return resp, nil
}
// FindOneWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodeLikeRelationsModelRepoImp) FindOneWithId(ctx rest.Context, id int64,) (*models.EpisodeLikeRelationsModel, error) { 
	resp := &models.EpisodeLikeRelationsModel{} 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("id = ?", id)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithEpisodeIdUserId 根据唯一索引 uniq_idx_episode_id_user_id 生成
func (repo *EpisodeLikeRelationsModelRepoImp) FindOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64,userId int64,) (*models.EpisodeLikeRelationsModel, error) { 
	resp := &models.EpisodeLikeRelationsModel{} 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("episode_id = ?", episodeId).
        Where("user_id = ?", userId)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodeLikeRelationsModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.EpisodeLikeRelationsModel) error { 
	resp := data 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("id = ?", id)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
} 

// FirstOrCreateWithEpisodeIdUserId 根据唯一索引 uniq_idx_episode_id_user_id 生成
func (repo *EpisodeLikeRelationsModelRepoImp) FirstOrCreateWithEpisodeIdUserId(ctx rest.Context, episodeId int64,userId int64, data *models.EpisodeLikeRelationsModel) error { 
	resp := data 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("episode_id = ?", episodeId).
        Where("user_id = ?", userId)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
} 

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodeLikeRelationsModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64,) error { 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("id = ?", id)
	if err := db.Delete(models.EpisodeLikeRelationsModel{}).Error;err!=nil{
		return errors.WithStack(err)
	}
	return nil
} 

// DeleteOneWithEpisodeIdUserId 根据唯一索引 uniq_idx_episode_id_user_id 生成
func (repo *EpisodeLikeRelationsModelRepoImp) DeleteOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64,userId int64,) error { 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("episode_id = ?", episodeId).
        Where("user_id = ?", userId)
	if err := db.Delete(models.EpisodeLikeRelationsModel{}).Error;err!=nil{
		return errors.WithStack(err)
	}
	return nil
} 

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodeLikeRelationsModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error { 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("id = ?", id)
	if err := db.Updates(attr).Error;err!=nil{
		return errors.WithStack(err)
	}
	return nil
} 

// UpdateOneWithEpisodeIdUserId 根据唯一索引 uniq_idx_episode_id_user_id 生成
func (repo *EpisodeLikeRelationsModelRepoImp) UpdateOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64,userId int64, attr map[string]interface{}) error { 
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
        Where("episode_id = ?", episodeId).
        Where("user_id = ?", userId)
	if err := db.Updates(attr).Error;err!=nil{
		return errors.WithStack(err)
	}
	return nil
} 

// Insert 默认生成的创建函数
func (repo *EpisodeLikeRelationsModelRepoImp) Insert(ctx rest.Context, data *models.EpisodeLikeRelationsModel) error {
	if err := repo.db.WithContext(ctx).Table("episode_like_relations").Create(data).Error;err!=nil{
		return errors.WithStack(err)
    }
	return nil
}
