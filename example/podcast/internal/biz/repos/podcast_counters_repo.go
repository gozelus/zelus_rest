package repos

import (
	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	models "github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
)

type PodcastCountersModelRepoImp struct {
	db db.MySQLDb
}

func NewPodcastCountersModelRepoImp(db db.MySQLDb) *PodcastCountersModelRepoImp {
	return &PodcastCountersModelRepoImp{db: db}
}

// FindManyWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PodcastCountersModelRepoImp) FindManyWithIdByTx(ctx rest.Context, tx db.MySQLDb, ids []int64) (map[int64]*models.PodcastCountersModel, error) {
	resp := map[int64]*models.PodcastCountersModel{}
	var results []*models.PodcastCountersModel
	db := tx.Table(ctx, "podcast_counters").
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
func (repo *PodcastCountersModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.PodcastCountersModel, error) {
	resp := map[int64]*models.PodcastCountersModel{}
	var results []*models.PodcastCountersModel
	db := repo.db.Table(ctx, "podcast_counters").
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
func (repo *PodcastCountersModelRepoImp) FindOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) (*models.PodcastCountersModel, error) {
	resp := &models.PodcastCountersModel{}
	db := tx.Table(ctx, "podcast_counters").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithId 根据唯一索引 PRIMARY 生成
func (repo *PodcastCountersModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.PodcastCountersModel, error) {
	resp := &models.PodcastCountersModel{}
	db := repo.db.Table(ctx, "podcast_counters").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PodcastCountersModelRepoImp) FirstOrCreateWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, data *models.PodcastCountersModel) error {
	resp := data
	db := tx.Table(ctx, "podcast_counters").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *PodcastCountersModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.PodcastCountersModel) error {
	resp := data
	db := repo.db.Table(ctx, "podcast_counters").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PodcastCountersModelRepoImp) DeleteOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) error {
	db := tx.Table(ctx, "podcast_counters").
		Where("id = ?", id)
	if err := db.Delete(models.PodcastCountersModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *PodcastCountersModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.Table(ctx, "podcast_counters").
		Where("id = ?", id)
	if err := db.Delete(models.PodcastCountersModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PodcastCountersModelRepoImp) UpdateOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, attr map[string]interface{}) error {
	db := tx.Table(ctx, "podcast_counters").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *PodcastCountersModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "podcast_counters").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// InsertByTx 默认生成的创建函数, 使用 tx 句柄
func (repo *PodcastCountersModelRepoImp) InsertByTx(ctx rest.Context, tx db.MySQLDb, data *models.PodcastCountersModel) error {
	if err := tx.Table(ctx, "podcast_counters").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *PodcastCountersModelRepoImp) Insert(ctx rest.Context, data *models.PodcastCountersModel) error {
	if err := repo.db.Table(ctx, "podcast_counters").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}