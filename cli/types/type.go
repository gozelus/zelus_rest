package types

type UserGetRequest struct {
	UserID int64
}
type UserGetResponse struct {
	UserID   int64
	UserName string
}
