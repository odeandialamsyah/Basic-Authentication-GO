package middleware

import (
	"net/http"

	"basic_authentication_go/config"
	"basic_authentication_go/models"
	"basic_authentication_go/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := utils.VerifyJWT(tokenString)
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Next()
	}
}

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(uint)

		var user models.User
		if err := config.DB.Preload("Role").First(&user, userID).Error; err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if user.Role.Name != requiredRole {
			c.JSON(http.StatusForbidden, gin. H {"error": "Access Denied"})
		}

		c.Next()
	}
}