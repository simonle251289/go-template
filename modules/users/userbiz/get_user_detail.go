package userbiz

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"template/components/appcontext"
	"template/modules/users/usermodels"
	"template/modules/users/userrepos"
	"template/utils/app_errors"
)

type GetUserDetailRepos interface {
	FindById(userId string) (*usermodels.UserResponse, error)
	FindByUsername(username string) (*usermodels.UserEntity, error)
}

type getUserDetailBiz struct {
	repos GetUserDetailRepos
}

func NewGetUserDetailBiz(ctx appcontext.AppContext) *getUserDetailBiz {
	repos := userrepos.NewUserRepos(ctx)
	return &getUserDetailBiz{
		repos: repos,
	}
}

func (biz *getUserDetailBiz) GetUserDetailById(userId string) (*usermodels.UserResponse, error) {
	user, err := biz.repos.FindById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app_errors.NewError(err, http.StatusNotFound, app_errors.ItemNotFound)
	} else if err != nil {
		return nil, app_errors.NewError(err, http.StatusNotFound, app_errors.ItemNotFound)
	}
	return user, nil
}

func (biz *getUserDetailBiz) GetUserDetailByUsername(username string) (*usermodels.UserEntity, error) {
	user, err := biz.repos.FindByUsername(username)
	return user, err
}
