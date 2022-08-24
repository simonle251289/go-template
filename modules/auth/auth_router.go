package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/components/appcontext"
	"template/modules/auth/authhandlers"
	"template/utils/app_errors"
)

func RegisterAuthRoute(ctx appcontext.AppContext, engine *gin.Engine) {
	auth := engine.Group("/auth")
	{
		auth.POST("/login", authhandlers.Login(ctx))
		auth.POST("/register", func(context *gin.Context) {
			panic(app_errors.NewError(nil, http.StatusBadRequest, app_errors.InvalidAccess))
		})
	}
}
