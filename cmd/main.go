package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
