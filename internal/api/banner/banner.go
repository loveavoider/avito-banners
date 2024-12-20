package banner

import (
	"errors"
	"github.com/loveavoider/avito-banners/internal/api/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/converter/banner"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/service"
	"github.com/loveavoider/avito-banners/internal/validator"
)

// TODO накидать валидаторы куда не хватает, чтобы всё одинаково выглядело

type BannerController struct {
	bannerService   service.BannerService
	bannerValidator validator.BannerValidator
}

func (bc BannerController) GetBanners(c *gin.Context) {
	getBanners, err := banner.GetBanners(c)

	if err != nil {
		util.WriteError(c, http.StatusBadRequest, err)
		return
	}

	banners, err := bc.bannerService.GetBanners(*getBanners)

	if err != nil {
		if errors.Is(err, service.BannersNotFound) {
			util.WriteError(c, http.StatusNotFound, err)
			return
		}

		util.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, banners)
}

func (bc BannerController) GetUserBanner(c *gin.Context) {
	getUserBanner, err := banner.GetUserBanner(c)

	if err != nil {
		util.WriteError(c, http.StatusBadRequest, err)
		return
	}

	content, err := bc.bannerService.GetUserBanner(*getUserBanner)

	if err != nil {

		if errors.Is(err, service.BannersNotFound) {
			util.WriteError(c, http.StatusNotFound, err)
			return
		}

		util.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, content)
}

func (bc BannerController) CreateBanner(c *gin.Context) {

	modelBanner, err := bc.bannerValidator.ValidationCreateBanner(c)

	if err != nil {
		util.WriteError(c, http.StatusBadRequest, err)
		return
	}

	id, err := bc.bannerService.CreateBanner(*modelBanner)

	if err != nil {
		util.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"banner_id": id})
}

func (bc BannerController) UpdateBanner(c *gin.Context) {

	updateBanner, err := bc.bannerValidator.ValidationUpdateBanner(c)

	util.WriteError(c, http.StatusBadRequest, err)

	err = bc.bannerService.UpdateBanner(*updateBanner)

	if err != nil {
		if errors.Is(err, service.BannersNotFound) {
			util.WriteError(c, http.StatusNotFound, err)
			return
		}

		util.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "")
}

func (bc BannerController) DeleteBanner(c *gin.Context) {

	modelBanner := model.Banner{}
	err := c.BindUri(&modelBanner)

	util.WriteError(c, http.StatusBadRequest, err)

	err = bc.bannerService.DeleteBanner(modelBanner)

	if err != nil {

		if errors.Is(err, service.BannersNotFound) {
			util.WriteError(c, http.StatusNotFound, err)
			return
		}

		util.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusNoContent, "")
}

func NewController(bs service.BannerService, bv validator.BannerValidator) *BannerController {
	return &BannerController{
		bannerService:   bs,
		bannerValidator: bv,
	}
}
