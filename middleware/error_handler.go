package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/components/appcontext"
	"template/utils/app_errors"
)

func ErrorHandler(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.Header("Content-Type", "application/json")
				if appErr, ok := err.(*app_errors.AppError); ok {
					context.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}
				appErr := app_errors.NewError(err.(error), http.StatusInternalServerError, app_errors.Generic)
				context.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		context.Next()
	}
}
