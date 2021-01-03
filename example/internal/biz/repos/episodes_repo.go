package repos

import (
	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	models "github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
)

type EpisodesModelRepoImp struct {
	db db.MySQLDb
}

func NewEpisodesModelRepoImp(db db.MySQLDb) *EpisodesModelRepoImp {
	return &EpisodesModelRepoImp{db: db}
}

// FindManyWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *EpisodesModelRepoImp) FindManyWithIdByTx(ctx rest.Context, tx db.MySQLDb, ids []int64) (map[int64]*models.EpisodesModel, error) {
	resp := map[int64]*models.EpisodesModel{}
	var results []*models.EpisodesModel
	db := tx.Table(ctx, "episodes").
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
func (repo *EpisodesModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.EpisodesModel, error) {
	resp := map[int64]*models.EpisodesModel{}
	var results []*models.EpisodesModel
	db := repo.db.Table(ctx, "episodes").
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
func (repo *EpisodesModelRepoImp) FindOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) (*models.EpisodesModel, error) {
	resp := &models.EpisodesModel{}
	db := tx.Table(ctx, "episodes").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodesModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.EpisodesModel, error) {
	resp := &models.EpisodesModel{}
	db := repo.db.Table(ctx, "episodes").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *EpisodesModelRepoImp) FirstOrCreateWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, data *models.EpisodesModel) error {
	resp := data
	db := tx.Table(ctx, "episodes").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodesModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.EpisodesModel) error {
	resp := data
	db := repo.db.Table(ctx, "episodes").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *EpisodesModelRepoImp) DeleteOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) error {
	db := tx.Table(ctx, "episodes").
		Where("id = ?", id)
	if err := db.Delete(models.EpisodesModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodesModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.Table(ctx, "episodes").
		Where("id = ?", id)
	if err := db.Delete(models.EpisodesModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *EpisodesModelRepoImp) UpdateOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, attr map[string]interface{}) error {
	db := tx.Table(ctx, "episodes").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *EpisodesModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "episodes").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// InsertByTx 默认生成的创建函数, 使用 tx 句柄
func (repo *EpisodesModelRepoImp) InsertByTx(ctx rest.Context, tx db.MySQLDb, data *models.EpisodesModel) error {
	if err := tx.Table(ctx, "episodes").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *EpisodesModelRepoImp) Insert(ctx rest.Context, data *models.EpisodesModel) error {
	if err := repo.db.Table(ctx, "episodes").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
