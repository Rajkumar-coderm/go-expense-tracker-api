package main

import (
	"github.com/expense-tracker-api/config"
	"github.com/expense-tracker-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.LoadJWTSecret()
	config.ConnectDB()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	routes.SetupRoutes(r)

	r.Run(":8080")
}
