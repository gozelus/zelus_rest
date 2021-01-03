package repos

import (
	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	models "github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
)

type UsersModelRepoImp struct {
	db db.MySQLDb
}

func NewUsersModelRepoImp(db db.MySQLDb) *UsersModelRepoImp {
	return &UsersModelRepoImp{db: db}
}

// FindManyWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) FindManyWithIdByTx(ctx rest.Context, tx db.MySQLDb, ids []int64) (map[int64]*models.UsersModel, error) {
	resp := map[int64]*models.UsersModel{}
	var results []*models.UsersModel
	db := tx.Table(ctx, "users").
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
func (repo *UsersModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.UsersModel, error) {
	resp := map[int64]*models.UsersModel{}
	var results []*models.UsersModel
	db := repo.db.Table(ctx, "users").
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
func (repo *UsersModelRepoImp) FindOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) (*models.UsersModel, error) {
	resp := &models.UsersModel{}
	db := tx.Table(ctx, "users").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.UsersModel, error) {
	resp := &models.UsersModel{}
	db := repo.db.Table(ctx, "users").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) FirstOrCreateWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, data *models.UsersModel) error {
	resp := data
	db := tx.Table(ctx, "users").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.UsersModel) error {
	resp := data
	db := repo.db.Table(ctx, "users").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) DeleteOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) error {
	db := tx.Table(ctx, "users").
		Where("id = ?", id)
	if err := db.Delete(models.UsersModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.Table(ctx, "users").
		Where("id = ?", id)
	if err := db.Delete(models.UsersModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) UpdateOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, attr map[string]interface{}) error {
	db := tx.Table(ctx, "users").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "users").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// InsertByTx 默认生成的创建函数, 使用 tx 句柄
func (repo *UsersModelRepoImp) InsertByTx(ctx rest.Context, tx db.MySQLDb, data *models.UsersModel) error {
	if err := tx.Table(ctx, "users").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *UsersModelRepoImp) Insert(ctx rest.Context, data *models.UsersModel) error {
	if err := repo.db.Table(ctx, "users").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
