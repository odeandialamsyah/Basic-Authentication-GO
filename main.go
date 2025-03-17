package main

import (
	"basic_authentication_go/config"
	"basic_authentication_go/models"
	"basic_authentication_go/routes"
	"basic_authentication_go/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	models.MigrateUsers(config.DB)
	utils.SeedRole(config.DB)

	r := gin.Default()

	routes.AuthRoutes(r, config.DB)

	r.Run(":8080")
}