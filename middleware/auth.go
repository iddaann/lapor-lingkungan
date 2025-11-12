package middleware

import (
	"laporan-lingkungan/config"
	"laporan-lingkungan/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", uint(claims["user_id"].(float64)))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}
	}
}

func GetUserID(c *gin.Context) uint {
	id, _ := c.Get("user_id")
	return id.(uint)
}

func AutoMigrate() {
	config.DB.AutoMigrate(&models.User{}, &models.Report{})
}
