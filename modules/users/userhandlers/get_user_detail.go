package userhandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/components/appcontext"
	"template/modules/users/userbiz"
	"template/utils/app_errors"
	"template/utils/dataresponse"
)

func GetUserDetail(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.Param("userId")
		if len(userId) == 0 {
			panic(app_errors.NewError(nil, http.StatusBadRequest, app_errors.MissingRequiredField))
			return
		}
		biz := userbiz.NewGetUserDetailBiz(ctx)
		user, err := biz.GetUserDetailById(userId)
		if err != nil {
			panic(err)
			return
		}
		context.JSON(http.StatusOK, dataresponse.NewSuccessResponse(user))
	}
}
