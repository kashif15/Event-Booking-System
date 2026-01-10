package middleware

import (
	"net/http"
	"strings"

	"event-booking-api/internal/auth"
	"github.com/gin-gonic/gin"
)


func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Not authorized",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.ValidateToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Not authorized",
			})
			return
		}

		userId := int64(claims["user_id"].(float64))
		role := claims["role"].(string)

		c.Set("userId" , userId)
		c.Set("role", role)
		c.Next()
	}	
}