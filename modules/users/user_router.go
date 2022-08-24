package users

import (
	"github.com/gin-gonic/gin"
	"template/components/appcontext"
	"template/modules/users/userhandlers"
)

func RegisterUserRouter(ctx appcontext.AppContext, engine *gin.Engine) {
	auth := engine.Group("/users")
	{
		auth.GET("/:userId", userhandlers.GetUserDetail(ctx))
	}
}
