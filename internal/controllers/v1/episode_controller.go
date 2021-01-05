package v1_controllers

import (
	rest "github.com/gozelus/zelus_rest"

	api "github.com/gozelus/zelus_rest/internal"

	v1_services "github.com/gozelus/zelus_rest/internal/services/v1"
)

type EpisodeController struct {
	service *v1_services.EpisodeService
}

func NewEpisodeController(service *v1_services.EpisodeService) *EpisodeController {
	return &EpisodeController{service: service}
}

func (c *EpisodeController) GetEpisodeDetail(ctx rest.Context) {
	res := &api.EpisodeDetailResponse{}
	req := &api.EpisodeDetailRequest{}
	var err error
	if err := ctx.JSONQueryBind(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
		return
	}
	if res, err = c.service.GetEpisodeDetail(ctx, req); err != nil {
		ctx.RenderErrorJSON(nil, err)
		return
	}
	ctx.RenderOkJSON(res)
}
