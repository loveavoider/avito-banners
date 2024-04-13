package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("aud")

		if !ok {
			c.String(http.StatusUnauthorized, "")
			c.Abort()
			return
		}
		
		if role != "admin" {
			c.String(http.StatusForbidden, "")
			c.Abort()
			return
		}

		c.Next()
	}
}