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

func NewLoginBiz(repos LoginRepos) *loginBiz {
	return &loginBiz{
		repos: repos,
	}
}

func UserLogin(ctx appcontext.AppContext, userName string, password string) (*usermodels.UserEntity, error) {
	return authrepos.Login(ctx, userName, password)
}
