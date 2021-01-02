package repos

import (
	"time"

	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	models "github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
)

type SubPodcastUserRelationsModelRepoImp struct {
	db db.MySQLDb
}

// ListPodcastIdByUserIdOrderByCreateTime 根据索引 idx_user_id_create_time_podcast_id 生成
func (repo *SubPodcastUserRelationsModelRepoImp) ListPodcastIdByUserIdOrderByCreateTime(ctx rest.Context, userId int64, limit int64, createTime time.Time) ([]*models.SubPodcastUserRelationsModel, bool, error) {
	var resp []*models.SubPodcastUserRelationsModel
	var hasMore bool
	if err := repo.db.Table(ctx, "sub_podcast_user_relations").
		Select("podcast_id").
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
func (repo *SubPodcastUserRelationsModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.SubPodcastUserRelationsModel, error) {
	resp := map[int64]*models.SubPodcastUserRelationsModel{}
	var results []*models.SubPodcastUserRelationsModel
	db := repo.db.Table(ctx, "sub_podcast_user_relations").
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
func (repo *SubPodcastUserRelationsModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.SubPodcastUserRelationsModel, error) {
	resp := &models.SubPodcastUserRelationsModel{}
	db := repo.db.Table(ctx, "sub_podcast_user_relations").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithPodcastIdUserId 根据唯一索引 uniq_idx_podcast_id_user_id 生成
func (repo *SubPodcastUserRelationsModelRepoImp) FindOneWithPodcastIdUserId(ctx rest.Context, podcastId int64, userId int64) (*models.SubPodcastUserRelationsModel, error) {
	resp := &models.SubPodcastUserRelationsModel{}
	db := repo.db.Table(ctx, "sub_podcast_user_relations").
		Where("podcast_id = ?", podcastId).
		Where("user_id = ?", userId)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrInsertWithId 根据唯一索引 PRIMARY 生成
func (repo *SubPodcastUserRelationsModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.SubPodcastUserRelationsModel) error {
	resp := data
	db := repo.db.Table(ctx, "sub_podcast_user_relations").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrInsertWithPodcastIdUserId 根据唯一索引 uniq_idx_podcast_id_user_id 生成
func (repo *SubPodcastUserRelationsModelRepoImp) FirstOrCreateWithPodcastIdUserId(ctx rest.Context, podcastId int64, userId int64, data *models.SubPodcastUserRelationsModel) error {
	resp := data
	db := repo.db.Table(ctx, "sub_podcast_user_relations").
		Where("podcast_id = ?", podcastId).
		Where("user_id = ?", userId)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *SubPodcastUserRelationsModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.Table(ctx, "sub_podcast_user_relations").
		Where("id = ?", id)
	if err := db.Delete(models.SubPodcastUserRelationsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithPodcastIdUserId 根据唯一索引 uniq_idx_podcast_id_user_id 生成
func (repo *SubPodcastUserRelationsModelRepoImp) DeleteOneWithPodcastIdUserId(ctx rest.Context, podcastId int64, userId int64) error {
	db := repo.db.Table(ctx, "sub_podcast_user_relations").
		Where("podcast_id = ?", podcastId).
		Where("user_id = ?", userId)
	if err := db.Delete(models.SubPodcastUserRelationsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *SubPodcastUserRelationsModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "sub_podcast_user_relations").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithPodcastIdUserId 根据唯一索引 uniq_idx_podcast_id_user_id 生成
func (repo *SubPodcastUserRelationsModelRepoImp) UpdateOneWithPodcastIdUserId(ctx rest.Context, podcastId int64, userId int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "sub_podcast_user_relations").
		Where("podcast_id = ?", podcastId).
		Where("user_id = ?", userId)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *SubPodcastUserRelationsModelRepoImp) Insert(ctx rest.Context, data *models.SubPodcastUserRelationsModel) error {
	if err := repo.db.Table(ctx, "sub_podcast_user_relations").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
