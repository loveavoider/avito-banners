package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/service"
)

func TokenValidator(service service.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			c.String(http.StatusUnauthorized, "")
			c.Abort()
			return
		}

		splits := strings.Split(authorization, " ")

		if len(splits) != 2 {
			c.String(http.StatusUnauthorized, "")
			c.Abort()
			return
		}

		claims, err := service.Validate(splits[1])
		
		if err != nil {
			c.String(http.StatusUnauthorized, "")
			return
		}

		aud, jwtErr := claims.GetAudience()
		
		if jwtErr != nil {
			c.String(http.StatusUnauthorized, "")
			c.Abort()
			return
		}

		

		c.Set("aud", aud[0])
		c.Next()
	}
}