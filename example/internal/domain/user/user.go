package user

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/adapter/controllers"
	"github.com/gozelus/zelus_rest/example/internal/domain/user/entity"
	"github.com/gozelus/zelus_rest/example/internal/domain/user/po"
	"time"
)

type Repo interface {
	New(ctx rest.Context, userEntity *entity.UserEntity) error
	Save(ctx rest.Context, userEntity *entity.UserEntity)
}
type Domain struct {
	repo Repo
}

func NewDomain(repo Repo) *Domain {
	return &Domain{
		repo: repo,
	}
}

func (d *Domain) Register(ctx rest.Context, req *controllers.RegisterRequest) error {
	userEntity := entity.UserEntity{
		UserInfoPo: po.UserInfoPo{
			Nickname:  req.Nickname,
			Avatar:    req.Avatar,
			Age:       req.Age,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return d.repo.New(ctx, &userEntity)
}

func (Domain) Info(ctx rest.Context, req *controllers.InfoRequest) error {
	panic("implement me")
}

var _ controllers.UserDomain = &Domain{}
