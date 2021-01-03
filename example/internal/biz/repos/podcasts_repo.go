package repos

import (
	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	models "github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
)

type PodcastsModelRepoImp struct {
	db db.MySQLDb
}

func NewPodcastsModelRepoImp(db db.MySQLDb) *PodcastsModelRepoImp {
	return &PodcastsModelRepoImp{db: db}
}

// FindManyWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PodcastsModelRepoImp) FindManyWithIdByTx(ctx rest.Context, tx db.MySQLDb, ids []int64) (map[int64]*models.PodcastsModel, error) {
	resp := map[int64]*models.PodcastsModel{}
	var results []*models.PodcastsModel
	db := tx.Table(ctx, "podcasts").
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
func (repo *PodcastsModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.PodcastsModel, error) {
	resp := map[int64]*models.PodcastsModel{}
	var results []*models.PodcastsModel
	db := repo.db.Table(ctx, "podcasts").
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
func (repo *PodcastsModelRepoImp) FindOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) (*models.PodcastsModel, error) {
	resp := &models.PodcastsModel{}
	db := tx.Table(ctx, "podcasts").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithId 根据唯一索引 PRIMARY 生成
func (repo *PodcastsModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.PodcastsModel, error) {
	resp := &models.PodcastsModel{}
	db := repo.db.Table(ctx, "podcasts").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PodcastsModelRepoImp) FirstOrCreateWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, data *models.PodcastsModel) error {
	resp := data
	db := tx.Table(ctx, "podcasts").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *PodcastsModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.PodcastsModel) error {
	resp := data
	db := repo.db.Table(ctx, "podcasts").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PodcastsModelRepoImp) DeleteOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) error {
	db := tx.Table(ctx, "podcasts").
		Where("id = ?", id)
	if err := db.Delete(models.PodcastsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *PodcastsModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.Table(ctx, "podcasts").
		Where("id = ?", id)
	if err := db.Delete(models.PodcastsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PodcastsModelRepoImp) UpdateOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, attr map[string]interface{}) error {
	db := tx.Table(ctx, "podcasts").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *PodcastsModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "podcasts").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// InsertByTx 默认生成的创建函数, 使用 tx 句柄
func (repo *PodcastsModelRepoImp) InsertByTx(ctx rest.Context, tx db.MySQLDb, data *models.PodcastsModel) error {
	if err := tx.Table(ctx, "podcasts").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *PodcastsModelRepoImp) Insert(ctx rest.Context, data *models.PodcastsModel) error {
	if err := repo.db.Table(ctx, "podcasts").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
