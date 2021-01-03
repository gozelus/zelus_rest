package repos

import (
	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	models "github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
)

type UserBindsModelRepoImp struct {
	db db.MySQLDb
}

func NewUserBindsModelRepoImp(db db.MySQLDb) *UserBindsModelRepoImp {
	return &UserBindsModelRepoImp{db: db}
}

// FindManyWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) FindManyWithIdByTx(ctx rest.Context, tx db.MySQLDb, ids []int64) (map[int64]*models.UserBindsModel, error) {
	resp := map[int64]*models.UserBindsModel{}
	var results []*models.UserBindsModel
	db := tx.Table(ctx, "user_binds").
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
func (repo *UserBindsModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.UserBindsModel, error) {
	resp := map[int64]*models.UserBindsModel{}
	var results []*models.UserBindsModel
	db := repo.db.Table(ctx, "user_binds").
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
func (repo *UserBindsModelRepoImp) FindOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) (*models.UserBindsModel, error) {
	resp := &models.UserBindsModel{}
	db := tx.Table(ctx, "user_binds").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.UserBindsModel, error) {
	resp := &models.UserBindsModel{}
	db := repo.db.Table(ctx, "user_binds").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithBindCodeBindTypeByTx 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) FindOneWithBindCodeBindTypeByTx(ctx rest.Context, tx db.MySQLDb, bindCode string, bindType int64) (*models.UserBindsModel, error) {
	resp := &models.UserBindsModel{}
	db := tx.Table(ctx, "user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithBindCodeBindType 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) FindOneWithBindCodeBindType(ctx rest.Context, bindCode string, bindType int64) (*models.UserBindsModel, error) {
	resp := &models.UserBindsModel{}
	db := repo.db.Table(ctx, "user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) FirstOrCreateWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, data *models.UserBindsModel) error {
	resp := data
	db := tx.Table(ctx, "user_binds").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.UserBindsModel) error {
	resp := data
	db := repo.db.Table(ctx, "user_binds").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithBindCodeBindTypeByTx 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) FirstOrCreateWithBindCodeBindTypeByTx(ctx rest.Context, tx db.MySQLDb, bindCode string, bindType int64, data *models.UserBindsModel) error {
	resp := data
	db := tx.Table(ctx, "user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithBindCodeBindType 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) FirstOrCreateWithBindCodeBindType(ctx rest.Context, bindCode string, bindType int64, data *models.UserBindsModel) error {
	resp := data
	db := repo.db.Table(ctx, "user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) DeleteOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) error {
	db := tx.Table(ctx, "user_binds").
		Where("id = ?", id)
	if err := db.Delete(models.UserBindsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.Table(ctx, "user_binds").
		Where("id = ?", id)
	if err := db.Delete(models.UserBindsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithBindCodeBindTypeByTx 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) DeleteOneWithBindCodeBindTypeByTx(ctx rest.Context, tx db.MySQLDb, bindCode string, bindType int64) error {
	db := tx.Table(ctx, "user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.Delete(models.UserBindsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithBindCodeBindType 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) DeleteOneWithBindCodeBindType(ctx rest.Context, bindCode string, bindType int64) error {
	db := repo.db.Table(ctx, "user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.Delete(models.UserBindsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) UpdateOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, attr map[string]interface{}) error {
	db := tx.Table(ctx, "user_binds").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "user_binds").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithBindCodeBindTypeByTx 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) UpdateOneWithBindCodeBindTypeByTx(ctx rest.Context, tx db.MySQLDb, bindCode string, bindType int64, attr map[string]interface{}) error {
	db := tx.Table(ctx, "user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithBindCodeBindType 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) UpdateOneWithBindCodeBindType(ctx rest.Context, bindCode string, bindType int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// InsertByTx 默认生成的创建函数, 使用 tx 句柄
func (repo *UserBindsModelRepoImp) InsertByTx(ctx rest.Context, tx db.MySQLDb, data *models.UserBindsModel) error {
	if err := tx.Table(ctx, "user_binds").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *UserBindsModelRepoImp) Insert(ctx rest.Context, data *models.UserBindsModel) error {
	if err := repo.db.Table(ctx, "user_binds").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
