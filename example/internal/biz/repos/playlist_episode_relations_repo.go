package repos

import (
	"time"

	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	models "github.com/gozelus/zelus_rest/example/internal/data/db"
	"github.com/pkg/errors"
)

type PlaylistEpisodeRelationsModelRepoImp struct {
	db db.MySQLDb
}

func NewPlaylistEpisodeRelationsModelRepoImp(db db.MySQLDb) *PlaylistEpisodeRelationsModelRepoImp {
	return &PlaylistEpisodeRelationsModelRepoImp{db: db}
}

// ListEpisodeIdByPlaylistIdOrderByCreateTimeByTx 根据索引 idx_playlist_id_create_time_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) ListEpisodeIdByPlaylistIdOrderByCreateTimeByTx(ctx rest.Context, tx db.MySQLDb, playlistId int64, limit int64, createTime time.Time) ([]*models.PlaylistEpisodeRelationsModel, bool, error) {
	var resp []*models.PlaylistEpisodeRelationsModel
	var hasMore bool
	if err := tx.Table(ctx, "playlist_episode_relations").
		Select("episode_id").
		Where("playlist_id = ?", playlistId).
		Where("create_time < ?", createTime).
		Order("create_time desc").
		Limit(int(limit + 1)).
		Find(&resp); err != nil {
		return nil, false, errors.WithStack(err)
	}
	hasMore = len(resp) > int(limit)
	if hasMore {
		resp = resp[:len(resp)-1]
	}
	return resp, hasMore, nil
}

// ListEpisodeIdByPlaylistIdOrderByCreateTime 根据索引 idx_playlist_id_create_time_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) ListEpisodeIdByPlaylistIdOrderByCreateTime(ctx rest.Context, playlistId int64, limit int64, createTime time.Time) ([]*models.PlaylistEpisodeRelationsModel, bool, error) {
	var resp []*models.PlaylistEpisodeRelationsModel
	var hasMore bool
	if err := repo.db.Table(ctx, "playlist_episode_relations").
		Select("episode_id").
		Where("playlist_id = ?", playlistId).
		Where("create_time < ?", createTime).
		Order("create_time desc").
		Limit(int(limit + 1)).
		Find(&resp); err != nil {
		return nil, false, errors.WithStack(err)
	}
	hasMore = len(resp) > int(limit)
	if hasMore {
		resp = resp[:len(resp)-1]
	}
	return resp, hasMore, nil
}

// FindManyWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) FindManyWithIdByTx(ctx rest.Context, tx db.MySQLDb, ids []int64) (map[int64]*models.PlaylistEpisodeRelationsModel, error) {
	resp := map[int64]*models.PlaylistEpisodeRelationsModel{}
	var results []*models.PlaylistEpisodeRelationsModel
	db := tx.Table(ctx, "playlist_episode_relations").
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
func (repo *PlaylistEpisodeRelationsModelRepoImp) FindManyWithId(ctx rest.Context, ids []int64) (map[int64]*models.PlaylistEpisodeRelationsModel, error) {
	resp := map[int64]*models.PlaylistEpisodeRelationsModel{}
	var results []*models.PlaylistEpisodeRelationsModel
	db := repo.db.Table(ctx, "playlist_episode_relations").
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
func (repo *PlaylistEpisodeRelationsModelRepoImp) FindOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) (*models.PlaylistEpisodeRelationsModel, error) {
	resp := &models.PlaylistEpisodeRelationsModel{}
	db := tx.Table(ctx, "playlist_episode_relations").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithId 根据唯一索引 PRIMARY 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) FindOneWithId(ctx rest.Context, id int64) (*models.PlaylistEpisodeRelationsModel, error) {
	resp := &models.PlaylistEpisodeRelationsModel{}
	db := repo.db.Table(ctx, "playlist_episode_relations").
		Where("id = ?", id)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithPlaylistIdEpisodeIdByTx 根据唯一索引 uniq_idx_playlist_id_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) FindOneWithPlaylistIdEpisodeIdByTx(ctx rest.Context, tx db.MySQLDb, playlistId int64, episodeId int64) (*models.PlaylistEpisodeRelationsModel, error) {
	resp := &models.PlaylistEpisodeRelationsModel{}
	db := tx.Table(ctx, "playlist_episode_relations").
		Where("playlist_id = ?", playlistId).
		Where("episode_id = ?", episodeId)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FindOneWithPlaylistIdEpisodeId 根据唯一索引 uniq_idx_playlist_id_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) FindOneWithPlaylistIdEpisodeId(ctx rest.Context, playlistId int64, episodeId int64) (*models.PlaylistEpisodeRelationsModel, error) {
	resp := &models.PlaylistEpisodeRelationsModel{}
	db := repo.db.Table(ctx, "playlist_episode_relations").
		Where("playlist_id = ?", playlistId).
		Where("episode_id = ?", episodeId)
	if err := db.First(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

// FirstOrCreateWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) FirstOrCreateWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, data *models.PlaylistEpisodeRelationsModel) error {
	resp := data
	db := tx.Table(ctx, "playlist_episode_relations").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithId 根据唯一索引 PRIMARY 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) FirstOrCreateWithId(ctx rest.Context, id int64, data *models.PlaylistEpisodeRelationsModel) error {
	resp := data
	db := repo.db.Table(ctx, "playlist_episode_relations").
		Where("id = ?", id)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithPlaylistIdEpisodeIdByTx 根据唯一索引 uniq_idx_playlist_id_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) FirstOrCreateWithPlaylistIdEpisodeIdByTx(ctx rest.Context, tx db.MySQLDb, playlistId int64, episodeId int64, data *models.PlaylistEpisodeRelationsModel) error {
	resp := data
	db := tx.Table(ctx, "playlist_episode_relations").
		Where("playlist_id = ?", playlistId).
		Where("episode_id = ?", episodeId)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// FirstOrCreateWithPlaylistIdEpisodeId 根据唯一索引 uniq_idx_playlist_id_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) FirstOrCreateWithPlaylistIdEpisodeId(ctx rest.Context, playlistId int64, episodeId int64, data *models.PlaylistEpisodeRelationsModel) error {
	resp := data
	db := repo.db.Table(ctx, "playlist_episode_relations").
		Where("playlist_id = ?", playlistId).
		Where("episode_id = ?", episodeId)
	if err := db.FirstOrCreate(resp); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) DeleteOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64) error {
	db := tx.Table(ctx, "playlist_episode_relations").
		Where("id = ?", id)
	if err := db.Delete(models.PlaylistEpisodeRelationsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithId 根据唯一索引 PRIMARY 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) DeleteOneWithId(ctx rest.Context, id int64) error {
	db := repo.db.Table(ctx, "playlist_episode_relations").
		Where("id = ?", id)
	if err := db.Delete(models.PlaylistEpisodeRelationsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithPlaylistIdEpisodeIdByTx 根据唯一索引 uniq_idx_playlist_id_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) DeleteOneWithPlaylistIdEpisodeIdByTx(ctx rest.Context, tx db.MySQLDb, playlistId int64, episodeId int64) error {
	db := tx.Table(ctx, "playlist_episode_relations").
		Where("playlist_id = ?", playlistId).
		Where("episode_id = ?", episodeId)
	if err := db.Delete(models.PlaylistEpisodeRelationsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DeleteOneWithPlaylistIdEpisodeId 根据唯一索引 uniq_idx_playlist_id_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) DeleteOneWithPlaylistIdEpisodeId(ctx rest.Context, playlistId int64, episodeId int64) error {
	db := repo.db.Table(ctx, "playlist_episode_relations").
		Where("playlist_id = ?", playlistId).
		Where("episode_id = ?", episodeId)
	if err := db.Delete(models.PlaylistEpisodeRelationsModel{}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithIdByTx 根据唯一索引 PRIMARY 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) UpdateOneWithIdByTx(ctx rest.Context, tx db.MySQLDb, id int64, attr map[string]interface{}) error {
	db := tx.Table(ctx, "playlist_episode_relations").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithId 根据唯一索引 PRIMARY 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) UpdateOneWithId(ctx rest.Context, id int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "playlist_episode_relations").
		Where("id = ?", id)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithPlaylistIdEpisodeIdByTx 根据唯一索引 uniq_idx_playlist_id_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) UpdateOneWithPlaylistIdEpisodeIdByTx(ctx rest.Context, tx db.MySQLDb, playlistId int64, episodeId int64, attr map[string]interface{}) error {
	db := tx.Table(ctx, "playlist_episode_relations").
		Where("playlist_id = ?", playlistId).
		Where("episode_id = ?", episodeId)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateOneWithPlaylistIdEpisodeId 根据唯一索引 uniq_idx_playlist_id_episode_id 生成
func (repo *PlaylistEpisodeRelationsModelRepoImp) UpdateOneWithPlaylistIdEpisodeId(ctx rest.Context, playlistId int64, episodeId int64, attr map[string]interface{}) error {
	db := repo.db.Table(ctx, "playlist_episode_relations").
		Where("playlist_id = ?", playlistId).
		Where("episode_id = ?", episodeId)
	if err := db.Updates(attr); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// InsertByTx 默认生成的创建函数, 使用 tx 句柄
func (repo *PlaylistEpisodeRelationsModelRepoImp) InsertByTx(ctx rest.Context, tx db.MySQLDb, data *models.PlaylistEpisodeRelationsModel) error {
	if err := tx.Table(ctx, "playlist_episode_relations").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Insert 默认生成的创建函数
func (repo *PlaylistEpisodeRelationsModelRepoImp) Insert(ctx rest.Context, data *models.PlaylistEpisodeRelationsModel) error {
	if err := repo.db.Table(ctx, "playlist_episode_relations").Insert(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
