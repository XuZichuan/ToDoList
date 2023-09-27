package router

import (
	"ToDoList/app/gateway/handler"
	"ToDoList/app/gateway/middleware"
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

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("tasks", handler.ListTaskHandler)
			authed.POST("task", handler.CreateTaskHandler)
			authed.GET("task/:id", handler.GetTaskHandler)       // task_id
			authed.PUT("task/:id", handler.UpdateTaskHandler)    // task_id
			authed.DELETE("task/:id", handler.DeleteTaskHandler) // task_id
		}
	}
	return ginRouter
}
