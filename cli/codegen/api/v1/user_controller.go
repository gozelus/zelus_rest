package v1
type UserService interface {
}
type UserController struct {
	service UserService
}
func NewUserController(service UserService) *UserController {
	return &UserController{service : service}
}
