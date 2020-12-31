package models

import (
	"github.com/gozelus/zelus_rest"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type EpisodeLikeRelationsRepo struct {
	db *gorm.DB
}

type EpisodeLikeRelationsModel struct {
	Id         int64     `gorm:"id"`
	EpisodeId  int64     `gorm:"episode_id"`
	UserId     int64     `gorm:"user_id"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
}

func (repo *EpisodeLikeRelationsRepo) Insert(ctx rest.Context, data *EpisodeLikeRelationsModel) error {
	if err := repo.db.WithContext(ctx).Table("episode_like_relations").Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *EpisodeLikeRelationsRepo) FirstOrCreateWithId(ctx rest.Context, id int64, data *EpisodeLikeRelationsModel) error {
	resp := data
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func (repo *EpisodeLikeRelationsRepo) FindOneWithId(ctx rest.Context, id int64) (*EpisodeLikeRelationsModel, error) {
	resp := &EpisodeLikeRelationsModel{}
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("id = ?", id)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}
func (repo *EpisodeLikeRelationsRepo) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("id = ?", id)
	if err := db.Updates(attr).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func (repo *EpisodeLikeRelationsRepo) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("id = ?", id)
	if err := db.Delete(EpisodeLikeRelationsModel{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// 非唯一索引

func (repo *EpisodeLikeRelationsRepo) FirstOrCreateWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64, data *EpisodeLikeRelationsModel) error {
	resp := data
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func (repo *EpisodeLikeRelationsRepo) FindOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64) (*EpisodeLikeRelationsModel, error) {
	resp := &EpisodeLikeRelationsModel{}
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}
func (repo *EpisodeLikeRelationsRepo) UpdateOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64, attr map[string]interface{}) error {
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.Updates(attr).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func (repo *EpisodeLikeRelationsRepo) DeleteOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64) error {
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.Delete(EpisodeLikeRelationsModel{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// 非唯一索引

// select last where first order by mid

func (repo *EpisodeLikeRelationsRepo) SelectEpisodeIdWithUserIdOrderByCreateTime(ctx rest.Context, userId int64, limit int, lastScore time.Time) ([]*EpisodeLikeRelationsModel, error) {
	var resp []*EpisodeLikeRelationsModel
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("user_id = ?", userId)
	if err := db.Order("create_time desc").Limit(limit).Find(&resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}
