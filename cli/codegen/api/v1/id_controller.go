package v1

type IdService interface {
}
type IdController struct {
	service IdService
}

func NewIdController(service IdService) *IdController {
	return &IdController{service: service}
}
