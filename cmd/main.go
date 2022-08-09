package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"template/components/appcontext"
	"template/configs"
	"template/middleware"
	"template/modules/auth"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Default())
	config := configs.LoadConfig()
	appCtx := appcontext.NewAppContext(&config)
	//Register route
	auth.RegisterAuthRoute(appCtx, r)
	if err := r.Run("127.0.0.1:" + strconv.Itoa(config.Port)); err != nil {
		log.Fatalln(err)
	}
}
