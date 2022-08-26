package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"template/components/appcontext"
	"template/configs"
	"template/middleware"
	"template/modules/auth"
	"template/modules/users"
	"template/utils/app_errors"
)

func main() {
	config := configs.LoadConfig()
	appCtx := appcontext.NewAppContext(&config)
	//
	r := gin.Default()
	r.Use(middleware.Default())
	r.Use(middleware.ErrorHandler(appCtx))
	BindRouter(appCtx, r)
	if err := r.Run(":" + strconv.Itoa(config.Port)); err != nil {
		log.Fatalln(err)
	}
}

func BindRouter(ctx appcontext.AppContext, engine *gin.Engine) {
	//Kubernetes health check
	engine.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	//Register route
	auth.RegisterAuthRoute(ctx, engine)
	users.RegisterUserRouter(ctx, engine)
	//define no route
	engine.NoRoute(func(ctx *gin.Context) {
		panic(app_errors.NewError(nil, http.StatusNotFound, app_errors.InvalidUrl))
	})
}
