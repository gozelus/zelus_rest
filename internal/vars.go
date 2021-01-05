package api

type UserInfo struct {
	Id        int64  `json:"user_id"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
}

type SendLoginOrRegisterRequest struct {
	Phone string `json:"phone"`
}

type SendLoginOrRegisterResponse struct {
}

type LoginByPhoneCodeRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type LoginByPhoneCodeResponse struct {
	UserInfo *UserInfo `json:"user_info"`
}

type EpisodeSummary struct {
	Id       int64  `json:"episode_id"`
	Title    string `json:"title"`
	Subtitle string `json:"sub_title"`
}

type EpisodeDetail struct {
}

type EpisodeDetailRequest struct {
	EpisodeId int64 `query:"episode_id" validate:"required"`
}

type EpisodeDetailResponse struct {
	Episode *EpisodeDetail `json:"episode"`
}
