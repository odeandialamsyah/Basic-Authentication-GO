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

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	routes.AuthRoutes(r, config.DB)

	r.Run(":8080")
}