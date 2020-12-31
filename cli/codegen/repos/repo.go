package repos

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/cli/codegen/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type EpisodeLikeRelationsModelRepoImp struct {
	db *gorm.DB
}

func (repo *EpisodeLikeRelationsModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.EpisodeLikeRelationsModel, error) {
	resp := map[int64]*models.EpisodeLikeRelationsModel{}
	var results []*models.EpisodeLikeRelationsModel
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("id in (?)", ids)
	if err := db.Find(&results).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	for _, r := range results {
		resp[r.Id] = r
	}
	return resp, nil
}
func (repo *EpisodeLikeRelationsModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.EpisodeLikeRelationsModel, error) {
	resp := &models.EpisodeLikeRelationsModel{}
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("id = ?", id)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

func (repo *EpisodeLikeRelationsModelRepoImp) FindOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64) (*models.EpisodeLikeRelationsModel, error) {
	resp := &models.EpisodeLikeRelationsModel{}
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.First(resp).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

func (repo *EpisodeLikeRelationsModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.EpisodeLikeRelationsModel) error {
	resp := data
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *EpisodeLikeRelationsModelRepoImp) FirstOrCreateWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64, data *models.EpisodeLikeRelationsModel) error {
	resp := data
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.FirstOrCreate(resp).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *EpisodeLikeRelationsModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("id = ?", id)
	if err := db.Delete(models.EpisodeLikeRelationsModel{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *EpisodeLikeRelationsModelRepoImp) DeleteOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64) error {
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.Delete(models.EpisodeLikeRelationsModel{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *EpisodeLikeRelationsModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("id = ?", id)
	if err := db.Updates(attr).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *EpisodeLikeRelationsModelRepoImp) UpdateOneWithEpisodeIdUserId(ctx rest.Context, episodeId int64, userId int64, attr map[string]interface{}) error {
	db := repo.db.WithContext(ctx).Table("episode_like_relations").
		Where("episode_id = ?", episodeId).
		Where("user_id = ?", userId)
	if err := db.Updates(attr).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *EpisodeLikeRelationsModelRepoImp) Insert(ctx rest.Context, data *models.EpisodeLikeRelationsModel) error {
	if err := repo.db.WithContext(ctx).Table("episode_like_relations").Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
