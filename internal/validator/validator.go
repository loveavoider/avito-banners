package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/model"
)

type BannerValidator interface {
	ValidationCreateBanner(*gin.Context) (*model.Banner, error)
	ValidationUpdateBanner(*gin.Context) (*model.UpdateBanner, error)
}
