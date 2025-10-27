package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/j0rqy/presentorAPI/health"
	"github.com/j0rqy/presentorAPI/user"
	"github.com/j0rqy/presentorAPI/userstore"
)

func Routes(router *gin.Engine) {
	userStore := userstore.NewPostgresStore()
	userService := user.NewUserService(userStore)
	userHandler := user.NewUserController(userService)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", health.GetHealthHandler)
			userRoute := v1.Group("/user")
			{
				userRoute.POST("", userHandler.CreateUserHandler)
			}
		}
	}
}
