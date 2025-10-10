package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/j0rqy/presentorAPI/internal/controllers"
)

func Routes(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", controller.Health)
		}
	}
}
