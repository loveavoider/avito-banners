package api

import "github.com/gin-gonic/gin"

type BannerController interface {
	CreateBanner(c *gin.Context)
	UpdateBanner(c *gin.Context)
	DeleteBanner(c *gin.Context)
	GetUserBanner(c *gin.Context)
	GetBanners(c *gin.Context)
}

type TokenController interface {
	GetToken(c *gin.Context)
}
