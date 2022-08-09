package authhandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/components/appcontext"
	"template/modules/auth/authbiz"
	"template/modules/auth/authmodels"
	"template/utils/dataresponse"
)

func Login(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data authmodels.LoginInput
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing require field",
			})
		}
		biz := authbiz.NewLoginBiz(ctx)
		user, err := biz.UserLogin(ctx, data.UserName, data.Password)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, dataresponse.NewSuccessResponse(user))
	}
}
