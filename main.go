package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/j0rqy/presentorAPI/database"
	"github.com/j0rqy/presentorAPI/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDB()
	defer database.CloseDB()
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
