package health

import "github.com/gin-gonic/gin"

func GetHealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
