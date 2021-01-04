package service

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/id_generator/api"
	v1 "github.com/gozelus/zelus_rest/example/id_generator/internal/controllers/v1"
	snowflake "github.com/gozelus/zelus_rest/example/id_generator/pkg"
)

type IdService struct {
	snowflake *snowflake.Node
}

func NewIdService() *IdService {
	return &IdService{
		snowflake: snowflake.SnowFlake,
	}
}
func (i *IdService) Single(ctx rest.Context, request *api.SingleIDGetRequest) (*api.SingleIDGetResponse, error) {
	res := &api.SingleIDGetResponse{}
	res.ID = i.snowflake.Generate()
	return res, nil
}

func (i *IdService) Batch(ctx rest.Context, request *api.BatchIDGetRequest) (*api.BatchIDGetResponse, error) {
	res := &api.BatchIDGetResponse{}
	res.IDs = i.snowflake.GenerateBatch(uint16(request.Count))
	return res, nil
}

var _ v1.IdService = &IdService{}
