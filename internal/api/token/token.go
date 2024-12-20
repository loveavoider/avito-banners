package token

import (
	"errors"
	"github.com/loveavoider/avito-banners/internal/api/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/service"
)

var incorrectRole = errors.New("incorrect role")

type TokenController struct {
	TokenService service.TokenService
}

func (tc TokenController) GetToken(c *gin.Context) {
	role := c.Query("role")
	if role == "user" || role == "admin" {
		token, err := tc.TokenService.Generate(role)

		if err != nil {
			util.WriteError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, token)
		return
	}

	util.WriteError(c, http.StatusBadRequest, incorrectRole)
}

func NewTockenController(tokenService service.TokenService) *TokenController {
	return &TokenController{
		TokenService: tokenService,
	}
}
