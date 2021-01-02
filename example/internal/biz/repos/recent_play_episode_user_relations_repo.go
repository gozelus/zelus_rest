package repos

import (
	"github.com/gozelus/zelus_rest/example/internal/data/db"
	"time"

	rest "github.com/gozelus/zelus_rest"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RecentPlayEpisodeUserRelationsModelRepoImp struct {
	db *gorm.DB
}

// ListEpisodeIdByUserIdOrderByCreateTime 根据索引 idx_user_id_target_type_create_time_target_id 生成
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) ListEpisodeIdByUserIdOrderByCreateTime(ctx rest.Context, userId int64, limit int64, createTime time.Time) ([]*models.RecentPlayEpisodeUserRelationsModel, bool, error) {
	var resp []*models.RecentPlayEpisodeUserRelationsModel
	var hasMore bool
	if err := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
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

// FindManyWithId 根据唯一索引 PRIMARY 生成
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.RecentPlayEpisodeUserRelationsModel, error) {
	resp := map[int64]*models.RecentPlayEpisodeUserRelationsModel{}
	var results []*models.RecentPlayEpisodeUserRelationsModel
	db := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
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
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.RecentPlayEpisodeUserRelationsModel, error) {
	resp := &models.RecentPlayEpisodeUserRelationsModel{}
	db := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
		Where("id = ?", id)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithUserIdEpisodeId 根据唯一索引 uniq_idx_user_id_target_id_target_type 生成
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) FindOneWithUserIdEpisodeId(ctx rest.Context, userId int64, episodeId int64) (*models.RecentPlayEpisodeUserRelationsModel, error) {
	resp := &models.RecentPlayEpisodeUserRelationsModel{}
	db := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
		Where("user_id = ?", userId).
		Where("episode_id = ?", episodeId)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.RecentPlayEpisodeUserRelationsModel) error {
	resp := data
	db := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithUserIdEpisodeId 根据唯一索引 uniq_idx_user_id_target_id_target_type 生成
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) FirstOrCreateWithUserIdEpisodeId(ctx rest.Context, userId int64, episodeId int64, data *models.RecentPlayEpisodeUserRelationsModel) error {
	resp := data
	db := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
		Where("user_id = ?", userId).
		Where("episode_id = ?", episodeId)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
		Where("id = ?", id)
	if err := db.Delete(models.RecentPlayEpisodeUserRelationsModel{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithUserIdEpisodeId 根据唯一索引 uniq_idx_user_id_target_id_target_type 生成
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) DeleteOneWithUserIdEpisodeId(ctx rest.Context, userId int64, episodeId int64) error {
	db := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
		Where("user_id = ?", userId).
		Where("episode_id = ?", episodeId)
	if err := db.Delete(models.RecentPlayEpisodeUserRelationsModel{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
		Where("id = ?", id)
	if err := db.Updates(attr).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithUserIdEpisodeId 根据唯一索引 uniq_idx_user_id_target_id_target_type 生成
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) UpdateOneWithUserIdEpisodeId(ctx rest.Context, userId int64, episodeId int64, attr map[string]interface{}) error {
	db := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").
		Where("user_id = ?", userId).
		Where("episode_id = ?", episodeId)
	if err := db.Updates(attr).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *RecentPlayEpisodeUserRelationsModelRepoImp) Insert(ctx rest.Context, data *models.RecentPlayEpisodeUserRelationsModel) error {
	if err := repo.db.WithContext(ctx).Table("recent_play_episode_user_relations").Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
