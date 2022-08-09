package auth

import (
	"github.com/gin-gonic/gin"
	"template/components/appcontext"
	"template/modules/auth/authhandlers"
)

func RegisterAuthRoute(ctx appcontext.AppContext, engine *gin.Engine) {
	auth := engine.Group("/auth")
	{
		auth.POST("/login", authhandlers.Login(ctx))
		auth.POST("/register", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
