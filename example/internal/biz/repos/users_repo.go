package repos

import (
	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UsersModelRepoImp struct {
	db *gorm.DB
}

// FindManyWithId 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.UsersModel, error) {
	resp := map[int64]*models.UsersModel{}
	var results []*models.UsersModel
	db := repo.db.WithContext(ctx).Table("users").
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
func (repo *UsersModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.UsersModel, error) {
	resp := &models.UsersModel{}
	db := repo.db.WithContext(ctx).Table("users").
		Where("id = ?", id)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.UsersModel) error {
	resp := data
	db := repo.db.WithContext(ctx).Table("users").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.WithContext(ctx).Table("users").
		Where("id = ?", id)
	if err := db.Delete(models.UsersModel{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *UsersModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.WithContext(ctx).Table("users").
		Where("id = ?", id)
	if err := db.Updates(attr).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *UsersModelRepoImp) Insert(ctx rest.Context, data *models.UsersModel) error {
	if err := repo.db.WithContext(ctx).Table("users").Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
