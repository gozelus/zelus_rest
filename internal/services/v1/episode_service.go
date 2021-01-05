package v1_services

import (
	"errors"

	rest "github.com/gozelus/zelus_rest"
	api "github.com/gozelus/zelus_rest/internal"
)

type EpisodeService struct {
	// 以后放入要依赖度的对象
}

func NewEpisodeService() *EpisodeService {
	return &EpisodeService{}
}

func (s *EpisodeService) GetEpisodeDetail(ctx rest.Context, request *api.EpisodeDetailRequest) (*api.EpisodeDetailResponse, error) {
	return nil, errors.New("no imp")
}
