package middleware

import (

	"myproject/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")

		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalid!"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalid!"})
		}

		claims, err := jwt.ParseToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid or expired"})
			c.Abort()
			return
		}

		c.Set("Username", claims.Username)
		c.Next()
	}
}
