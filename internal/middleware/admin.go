package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin only"})
			return
		}
		c.Next()
	}
}
