package repos

import (
	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	models "github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
)

type EpisodeCountersModelRepoImp struct {
	db db.MySQLDb
}

// FindManyWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) FindManyWithIdByTx(ctx rest.Context, tx db.MySQLDb, ids []int64) (map[int64]*models.EpisodeCountersModel, error) {
	resp := map[int64]*models.EpisodeCountersModel{}
	var results []*models.EpisodeCountersModel
	db := tx.Table(ctx, "episode_counters").
		Where("id in (?)", ids)
	if err := db.Find(&results); err != nil {
		return nil, errors.WithStack(err)
	}
	for _, r := range results {
		resp[r.Id] = r
	}
	return resp, nil
}

// FindManyWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.EpisodeCountersModel, error) {
	resp := map[int64]*models.EpisodeCountersModel{}
	var results []*models.EpisodeCountersModel
	db := repo.db.Table(ctx, "episode_counters").
		Where("id in (?)", ids)
	if err := db.Find(&results); err != nil {
		return nil, errors.WithStack(err)
	}
	for _, r := range results {
		resp[r.Id] = r
	}
	return resp, nil
}

// FindOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) FindOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) (*models.EpisodeCountersModel, error) {
	resp := &models.EpisodeCountersModel{}
	db := tx.Table(ctx, "episode_counters").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.EpisodeCountersModel, error) {
	resp := &models.EpisodeCountersModel{}
	db := repo.db.Table(ctx, "episode_counters").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) FirstOrCreateWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, data *models.EpisodeCountersModel) error {
	resp := data
	db := tx.Table(ctx, "episode_counters").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.EpisodeCountersModel) error {
	resp := data
	db := repo.db.Table(ctx, "episode_counters").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) DeleteOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) error {
	db := tx.Table(ctx, "episode_counters").
		Where("id = ?", id)
	if err := db.Delete(models.EpisodeCountersModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.Table(ctx, "episode_counters").
		Where("id = ?", id)
	if err := db.Delete(models.EpisodeCountersModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) UpdateOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, attr map[string]interface{}) error {
	db := tx.Table(ctx, "episode_counters").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodeCountersModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "episode_counters").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// InsertByTx 默认生成的创建函数, 使用 tx 句柄
func (repo *EpisodeCountersModelRepoImp) InsertByTx(ctx rest.Context, tx db.MySQLDb, data *models.EpisodeCountersModel) error {
	if err := tx.Table(ctx, "episode_counters").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *EpisodeCountersModelRepoImp) Insert(ctx rest.Context, data *models.EpisodeCountersModel) error {
	if err := repo.db.Table(ctx, "episode_counters").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
