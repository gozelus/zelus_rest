package v1

import (
	rest "github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/id_generator/api"
	apiErrors "github.com/gozelus/zelus_rest/example/id_generator/api/errors"
)

type IdService interface {
	Single(ctx rest.Context, request *api.SingleIDGetRequest) (*api.SingleIDGetResponse, error)
	Batch(ctx rest.Context, request *api.BatchIDGetRequest) (*api.BatchIDGetResponse, error)
}
type IdController struct {
	IdService IdService
}

func NewIdController(s IdService) *IdController {
	return &IdController{
		IdService: s,
	}
}

func (c *IdController) Batch(ctx rest.Context) {
	req := &api.BatchIDGetRequest{}
	res := &api.BatchIDGetResponse{}
	var err error
	if err := ctx.JSONQueryBind(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
		return
	}
	if res, err = c.IdService.Batch(ctx, req); err != nil {
		ctx.RenderErrorJSON(res, err)
		return
	}
	ctx.RenderOkJSON(res)
	return
}

func (c *IdController) Single(ctx rest.Context) {
	req := &api.SingleIDGetRequest{}
	res := &api.SingleIDGetResponse{}
	var err error
	if err := ctx.JSONQueryBind(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
	}
	if res, err = c.IdService.Single(ctx, req); err != nil {
		ctx.RenderErrorJSON(res, err)
		return
	}
	ctx.RenderOkJSON(res)
	return
}
