package util

import (
	"github.com/gin-gonic/gin"
)

func WriteError(c *gin.Context, code int, error error) {
	if len(error.Error()) == 0 {
		c.String(code, "")
	} else {
		c.JSON(
			code,
			gin.H{"error": error.Error()},
		)
	}
}
