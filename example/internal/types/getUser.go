package types

type GetUserRequest struct {
	UserID int64 `json:"user_id"`
}

type GetUserResponse struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}
