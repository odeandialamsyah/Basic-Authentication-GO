package routes

import (
    "auth-go/controllers"
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
}