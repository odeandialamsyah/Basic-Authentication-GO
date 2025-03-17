package routes

import (
	"basic_authentication_go/controllers"
	"basic_authentication_go/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.Engine, db *gorm.DB) {
    authController := controllers.AuthController{DB: db}

    authGroup := router.Group("/auth")
    {
        authGroup.POST("/register", authController.Register)
        authGroup.POST("/login", authController.Login)
    }

    adminGroup := router.Group("/auth/admin")
    adminGroup.Use(middleware.RoleMiddleware("admin")) 
    {
        adminGroup.GET("/dashboard", func (c *gin.Context)  {
            c.JSON(http.StatusOK, gin.H{"message": "Welcome to the admin dashboard"})
        })
    }
}