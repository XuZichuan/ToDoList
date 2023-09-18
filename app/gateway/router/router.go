package router

import (
	"ToDoList/app/gateway/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	ginRouter := gin.Default()
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "OK")
		})

		v1.POST("user/login", handler.UserLoginHandler)
		v1.POST("user/register", handler.UserRegisterHandler)
	}
	return ginRouter
}
