package types

type UserGetRequest struct {
	UserID   *int64  `json:"user_id" default:"123"`
	UserName string `json:"user_name",enum:"[1,b]"`
}
type UserGetResponse struct {
	UserID   int64  `json:"user_id" range:"[0, 10]"`
	UserName string `json:"user_name"`
}
type UserCreateRequest struct {
	UserID   int64
	UserName string
}
type UserCreateResponse struct {
	UserID   int64
	UserName string
}
