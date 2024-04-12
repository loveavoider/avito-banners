package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/service"
	"github.com/loveavoider/avito-banners/merror"
)

type TokenController struct {
	TokenService service.TokenService
}

func (tc TokenController) GetToken(c *gin.Context) {
	role := c.Query("role")
	if role == "user" || role == "admin" {
		token, err := tc.TokenService.Generate(role)

		if err != nil {
			c.JSON(http.StatusInternalServerError, &merror.MError{Message: err.Message})
			return
		}

		c.JSON(http.StatusOK, token)
		return
	}

	c.JSON(http.StatusBadRequest, &merror.MError{Message: "incorrect role"})
}

func NewTockenController(tokenService service.TokenService) *TokenController {
	return &TokenController{
		TokenService: tokenService,
	}
}

