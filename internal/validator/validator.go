package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/merror"
)

type BannerValidator interface {
	ValidationCreateBanner(*gin.Context) (*model.Banner, *merror.MError)
	ValidationUpdateBanner(*gin.Context) (*model.UpdateBanner, *merror.MError)
}