package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/j0rqy/presentorAPI/health"
	"github.com/j0rqy/presentorAPI/user"
	"github.com/j0rqy/presentorAPI/user_store"
)

func Routes(router *gin.Engine) {
	userStore := user_store.NewPostgresStore()
	userService := user.NewService(userStore)
	userHandler := user.NewHandler(userService)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", health.Health)
			userRoute := v1.Group("/user")
			{
				userRoute.POST("", userHandler.CreateUserEndPoint)
			}
		}
	}
}
