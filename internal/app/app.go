package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

type App struct {
	serviceProvider *serviceProvider

	router *gin.Engine
}

// TODO add init config and server and part routes
func NewApp() *App {
	a := &App{}

	a.initServiceProvider()
	a.initRouter()

	return a
}

func (a *App) initServiceProvider() {
	a.serviceProvider = newServiceProvider()
}

func (a *App) initRouter() {
	a.router = gin.Default()

	a.router.Use(gin.Recovery())
	a.router.Use(gin.Logger())

	bannerController := a.serviceProvider.BannerController()

	bannerGroup := a.router.Group("/banner").Use(a.serviceProvider.TokenValidator())
	bannerGroup.POST("", bannerController.CreateBanner)
	bannerGroup.PATCH("/:id", bannerController.UpdateBanner)
	bannerGroup.DELETE("/:id", bannerController.DeleteBanner)
	bannerGroup.GET("", bannerController.GetBanners)

	userBannerGroup := a.router.Group("/user_banner").Use(a.serviceProvider.TokenValidator())
	userBannerGroup.GET("", bannerController.GetUserBanner)

	tokenController := a.serviceProvider.TokenController()

	tokenGroup := a.router.Group("/token")
	tokenGroup.GET("", tokenController.GetToken)
}

func (a *App) Start() {
	a.router.Run(":" + os.Getenv("HOST_HTTP_PORT"))
}