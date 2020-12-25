package user

type DtoModel struct {
	Nickname string `json:"nickname"`
	UserID   int64  `json:"user_id"`
}
type RegisterRequest struct {
	Nickname string `json:"nickname",required:"true"`
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
