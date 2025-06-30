package main

import (
	"learn-go/backend-api/config"
	"learn-go/backend-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.InitDB()
	router :=gin.Default()
	router.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message" : "Hello World",
		})
	})

	router.Run(": " + config.GetEnv("APP_PORT", "3000"))
}