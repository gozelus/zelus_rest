package repos

import (
	"time"

	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	models "github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
)

type LikeEpisodeUserRelationsModelRepoImp struct {
	db db.MySQLDb
}

// ListEpisodeIdByUserIdOrderByCreateTime 根据索引 idx_user_id_create_time_episode_id 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) ListEpisodeIdByUserIdOrderByCreateTime(ctx rest.Context, userId int64, limit int64, createTime time.Time) ([]*models.LikeEpisodeUserRelationsModel, bool, error) {
	var resp []*models.LikeEpisodeUserRelationsModel
	var hasMore bool
	if err := repo.db.Table(ctx, "like_episode_user_relations").
		Select("episode_id").
		Where("user_id = ?", userId).
		Where("create_time < ?", createTime).
		Order("create_time desc").
		Limit(int(limit + 1)).
		Find(&resp); err != nil {
		return nil, false, errors.WithStack(err)
	}
	hasMore = len(resp) > int(limit)
	if hasMore {
		resp = resp[:len(resp)-1]
	}
	return resp, hasMore, nil
}

// FindManyWithId 根据唯一索引 PRIMARY 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.LikeEpisodeUserRelationsModel, error) {
	resp := map[int64]*models.LikeEpisodeUserRelationsModel{}
	var results []*models.LikeEpisodeUserRelationsModel
	db := repo.db.Table(ctx, "like_episode_user_relations").
		Where("id in (?)", ids)
	if err := db.Find(&results); err != nil {
		return nil, errors.WithStack(err)
	}
	for _, r := range results {
		resp[r.Id] = r
	}
	return resp, nil
}

// FindOneWithId 根据唯一索引 PRIMARY 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.LikeEpisodeUserRelationsModel, error) {
	resp := &models.LikeEpisodeUserRelationsModel{}
	db := repo.db.Table(ctx, "like_episode_user_relations").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithEpisodeIdUserId 根据唯一索引 uniq_idx_episode_id_user_id 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) FindOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64) (*models.LikeEpisodeUserRelationsModel, error) {
	resp := &models.LikeEpisodeUserRelationsModel{}
	db := repo.db.Table(ctx, "like_episode_user_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrInsertWithId 根据唯一索引 PRIMARY 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.LikeEpisodeUserRelationsModel) error {
	resp := data
	db := repo.db.Table(ctx, "like_episode_user_relations").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrInsertWithEpisodeIdUserId 根据唯一索引 uniq_idx_episode_id_user_id 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) FirstOrCreateWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64, data *models.LikeEpisodeUserRelationsModel) error {
	resp := data
	db := repo.db.Table(ctx, "like_episode_user_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.Table(ctx, "like_episode_user_relations").
		Where("id = ?", id)
	if err := db.Delete(models.LikeEpisodeUserRelationsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithEpisodeIdUserId 根据唯一索引 uniq_idx_episode_id_user_id 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) DeleteOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64) error {
	db := repo.db.Table(ctx, "like_episode_user_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.Delete(models.LikeEpisodeUserRelationsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "like_episode_user_relations").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithEpisodeIdUserId 根据唯一索引 uniq_idx_episode_id_user_id 生成
func (repo *LikeEpisodeUserRelationsModelRepoImp) UpdateOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "like_episode_user_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *LikeEpisodeUserRelationsModelRepoImp) Insert(ctx rest.Context, data *models.LikeEpisodeUserRelationsModel) error {
	if err := repo.db.Table(ctx, "like_episode_user_relations").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
