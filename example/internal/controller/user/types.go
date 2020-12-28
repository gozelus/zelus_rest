package user

type DtoModel struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Age      int    `json:"age"`
	UserID   int64  `json:"user_id"`
}
type RegisterRequest struct {
	Age      int    `json:"age" validate:"gte=18 lte=99 required"`
	Nickname string `json:"nickname" validate:"required"`
	Avatar   string `json:"avatar"`
}
type RegisterResponse struct {
	User *DtoModel `json:"user"`
}
type InfoRequest struct {
	UserID int64 `json:"user_id",range:"[1,2]"`
}
type InfoResponse struct {
	User *DtoModel `json:"user"`
}
