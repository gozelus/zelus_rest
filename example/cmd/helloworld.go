package main

import (
	"database/sql"
	"net/http"

	rest "github.com/gozelus/zelus_rest"
)

var (
	routes []rest.Route
)

func main() {
	s := rest.NewServer()
	routes = []rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/user/info",
			Handler: GetUser,
		},
	}
	s.AddRoute(routes...)
	s.Use(rest.LoggerMiddleware)
	s.Run("localhost", 8080)
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	gq := &getUserRequest{}
	if err := rest.JsonBodyFromRequest(req, gq); err != nil {
		rest.OkJson(w, map[string]interface{}{"message": gq.UserID})
	}
	res, err := GetUserWithService(gq)
	if err != nil {
		rest.Error(w, http.StatusInternalServerError, err)
		return
	}
	rest.OkJson(w, map[string]interface{}{"message": res.UserName})
}

type getUserResponse struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}
type getUserRequest struct {
	UserID int64 `json:"user_id"`
}

func GetUserWithService(getUserRequest *getUserRequest) (*getUserResponse, error) {
	nickname, err := GetUserWithDao(getUserRequest.UserID)
	if err != nil {
		return nil, err
	}
	return &getUserResponse{
		UserID:   getUserRequest.UserID,
		UserName: nickname,
	}, nil
}

func GetUserWithDao(userID int64) (string, error) {
	return "", sql.ErrNoRows
}
