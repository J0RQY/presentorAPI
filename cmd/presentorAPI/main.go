package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j0rqy/presentorAPI/internal/routes"
)

func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
