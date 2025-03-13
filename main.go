package main

import (
	"basic_authentication_go/config"
	"basic_authentication_go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	models.MigrateUsers(config.DB)

	r := gin.Default()

	routes.AuthRoutes(r, config.DB)

	r.Run(":8080")
}