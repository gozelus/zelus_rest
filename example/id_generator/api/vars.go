package api

type SingleIDGetRequest struct {
}
type SingleIDGetResponse struct {
	ID int64 `json:"id"`
}
type BatchIDGetRequest struct {
	Count int32 `json:"count,string"`
}
type BatchIDGetResponse struct {
	IDs []int64 `json:"ids"`
}
