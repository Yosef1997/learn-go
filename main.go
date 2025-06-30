package main

import (
	"learn-go/backend-api/config"
	"learn-go/backend-api/database"
	"learn-go/backend-api/routes"
)

func main() {
	config.LoadEnv()
	database.InitDB()
	r := routes.SetupRouter()
	

	r.Run(": " + config.GetEnv("APP_PORT", "3000"))
}