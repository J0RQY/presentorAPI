package routes

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"status": "ok",
				})
			})
		}
	}
}
