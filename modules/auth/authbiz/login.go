package authbiz

import (
	"template/components/appcontext"
	"template/modules/auth/authrepos"
	"template/modules/users/usermodels"
)

type LoginRepos interface {
	Login(ctx appcontext.AppContext, userName string, password string) (*usermodels.UserEntity, error)
}

type loginBiz struct {
	repos LoginRepos
}

func NewLoginBiz(ctx appcontext.AppContext) *loginBiz {
	repos := authrepos.NewUserFind(ctx)
	return &loginBiz{
		repos: repos,
	}
}

func (biz *loginBiz) UserLogin(ctx appcontext.AppContext, userName string, password string) (*usermodels.UserEntity, error) {
	return biz.repos.Login(ctx, userName, password)
}
