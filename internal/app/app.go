package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/api/middleware"
)

type App struct {
	serviceProvider *serviceProvider

	router *gin.Engine
}

func NewApp() *App {
	a := &App{}

	a.initServiceProvider()
	a.initRouter()

	return a
}

func (a *App) initServiceProvider() {
	a.serviceProvider = NewServiceProvider()
}

func (a *App) initRouter() {
	a.router = GetRouter(*a.serviceProvider)
}

func GetRouter(serviceProvider serviceProvider) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	bannerController := serviceProvider.BannerController()

	bannerGroup := router.Group("/banner").Use(serviceProvider.TokenValidator())

	bannerGroup.GET("", bannerController.GetBanners)

	bannerGroup.Use(middleware.ForAdmin())
	{
		bannerGroup.POST("", bannerController.CreateBanner)
		bannerGroup.PATCH("/:id", bannerController.UpdateBanner)
		bannerGroup.DELETE("/:id", bannerController.DeleteBanner)
	}

	userBannerGroup := router.Group("/user_banner").Use(serviceProvider.TokenValidator())
	userBannerGroup.GET("", bannerController.GetUserBanner)

	tokenController := serviceProvider.TokenController()

	tokenGroup := router.Group("/token")
	tokenGroup.GET("", tokenController.GetToken)

	return router
}

func (a *App) Start() {
	a.router.Run(":" + os.Getenv("HOST_HTTP_PORT"))
}
