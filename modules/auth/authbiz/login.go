package authbiz

import (
	"net/http"
	"template/components/appcontext"
	"template/modules/auth/authmodels"
	"template/modules/auth/authrepos"
	"template/modules/users/userbiz"
	"template/modules/users/usermodels"
	"template/utils"
	"template/utils/app_errors"
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

func (biz *loginBiz) UserLogin(ctx appcontext.AppContext, userName string, password string) (*authmodels.LoginResponse, error) {
	userBiz := userbiz.NewGetUserDetailBiz(ctx)
	user, err := userBiz.GetUserDetailByUsername(userName)
	if err != nil {
		return nil, app_errors.NewError(err, http.StatusBadRequest, app_errors.WrongUserNameOrPassword)
	}
	comparePw := utils.ValidateHash(password, user.Password)
	if comparePw == false {
		return nil, app_errors.NewError(err, http.StatusBadRequest, app_errors.WrongUserNameOrPassword)
	}
	userRes := user.ToResponse()
	var response = &authmodels.LoginResponse{
		User:  userRes,
		Token: utils.GenerateToken(ctx.GetConfig().JWT.AccessTTL, ctx.GetConfig().JWT.RefreshTTL, ctx.GetConfig().RSAPrivateKey, ctx.GetConfig().RSARefreshPrivateKey, &userRes),
	}
	return response, nil
}
