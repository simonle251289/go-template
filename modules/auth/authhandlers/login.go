package authhandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/components/appcontext"
	"template/modules/auth/authbiz"
	"template/modules/auth/authmodels"
	"template/utils/app_errors"
	"template/utils/dataresponse"
)

func Login(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data authmodels.LoginInput
		if err := context.ShouldBind(&data); err != nil {
			panic(app_errors.NewError(err, http.StatusBadRequest, app_errors.MissingRequiredField))
			return
		}

		if data.ValidateLoginInput() == false {
			panic(app_errors.NewError(nil, http.StatusBadRequest, app_errors.MissingRequiredField))
			return
		}

		biz := authbiz.NewLoginBiz(ctx)
		user, err := biz.UserLogin(ctx, data.UserName, data.Password)
		if err != nil {
			panic(err)
			return
		}
		context.JSON(http.StatusOK, dataresponse.NewSuccessResponse(user))
	}
}
