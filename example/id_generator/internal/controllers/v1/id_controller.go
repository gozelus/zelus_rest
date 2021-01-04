package v1

import (
	"github.com/gozelus/zelus_rest"
)

type IdService interface { 
	SingleGetID(ctx rest.Context, req *api.SingleIDGetRequest) (*api.SingleIDGetResponse, error) 
	BatchGetID(ctx rest.Context, req *api.BatchIDGetRequest) (*api.BatchIDGetResponse, error) 
}
type IdController struct {
	service IdService
}
func NewIdController(service IdService) *IdController {
	return &IdController{service : service}
}


func (c *IdController) SingleGetID(ctx rest.Context) {
	res := &api.SingleIDGetResponse{}
	req := &api.SingleIDGetRequest{}
	var err error 
	if err := ctx.JSONBodyBind(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
		return
	}
	if res, err = c.service.SingleGetID(ctx, req);err!=nil{
		ctx.RenderErrorJSON(nil, err)
		return
	}
	ctx.RenderOkJSON(res)
}

func (c *IdController) BatchGetID(ctx rest.Context) {
	res := &api.BatchIDGetResponse{}
	req := &api.BatchIDGetRequest{}
	var err error 
	if err := ctx.JSONBodyBind(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
		return
	}
	if res, err = c.service.BatchGetID(ctx, req);err!=nil{
		ctx.RenderErrorJSON(nil, err)
		return
	}
	ctx.RenderOkJSON(res)
}

