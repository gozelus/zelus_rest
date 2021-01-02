package repos

import (
	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserBindsModelRepoImp struct {
	db *gorm.DB
}

// FindManyWithId 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.UserBindsModel, error) {
	resp := map[int64]*models.UserBindsModel{}
	var results []*models.UserBindsModel
	db := repo.db.WithContext(ctx).Table("user_binds").
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
func (repo *UserBindsModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.UserBindsModel, error) {
	resp := &models.UserBindsModel{}
	db := repo.db.WithContext(ctx).Table("user_binds").
		Where("id = ?", id)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithBindCodeBindType 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) FindOneWithBindCodeBindType(ctx rest.Context, bindCode string, bindType int64) (*models.UserBindsModel, error) {
	resp := &models.UserBindsModel{}
	db := repo.db.WithContext(ctx).Table("user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.UserBindsModel) error {
	resp := data
	db := repo.db.WithContext(ctx).Table("user_binds").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithBindCodeBindType 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) FirstOrCreateWithBindCodeBindType(ctx rest.Context, bindCode string, bindType int64, data *models.UserBindsModel) error {
	resp := data
	db := repo.db.WithContext(ctx).Table("user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.WithContext(ctx).Table("user_binds").
		Where("id = ?", id)
	if err := db.Delete(models.UserBindsModel{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithBindCodeBindType 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) DeleteOneWithBindCodeBindType(ctx rest.Context, bindCode string, bindType int64) error {
	db := repo.db.WithContext(ctx).Table("user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.Delete(models.UserBindsModel{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UserBindsModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.WithContext(ctx).Table("user_binds").
		Where("id = ?", id)
	if err := db.Updates(attr).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithBindCodeBindType 根据唯一索引 uniq_idx_bind_code_bind_type 生成
func (repo *UserBindsModelRepoImp) UpdateOneWithBindCodeBindType(ctx rest.Context, bindCode string, bindType int64, attr map[string]interface{}) error {
	db := repo.db.WithContext(ctx).Table("user_binds").
		Where("bind_code = ?", bindCode).
		Where("bind_type = ?", bindType)
	if err := db.Updates(attr).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *UserBindsModelRepoImp) Insert(ctx rest.Context, data *models.UserBindsModel) error {
	if err := repo.db.WithContext(ctx).Table("user_binds").Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
