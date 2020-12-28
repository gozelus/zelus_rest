package user

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/adapter/controllers"
	"github.com/gozelus/zelus_rest/example/internal/domain/user/entity"
)

type Repo interface {
	New(ctx rest.Context) *entity.UserEntity
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

func (Domain) Register(ctx rest.Context, req *controllers.RegisterRequest) error {
	panic("implement me")
}

func (Domain) Info(ctx rest.Context, req *controllers.InfoRequest) error {
	panic("implement me")
}

var _ controllers.UserDomain = &Domain{}
