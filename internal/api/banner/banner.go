package banner

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/converter/banner"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/service"
	"github.com/loveavoider/avito-banners/internal/validator"
)

type BannerController struct {
	bannerService service.BannerService
	bannerValidator validator.BannerValidator
}

func (bc BannerController) GetBanners(c *gin.Context) {
	getBanners, err := banner.GetBanners(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Message},
		)
		return
	}

	banners, err := bc.bannerService.GetBanners(*getBanners)
	
	if err != nil {
		if err.Status == 404 {
			c.String(http.StatusNotFound, "")
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Message},
		)
		return
	}

	c.JSON(http.StatusOK, banners)
}

func (bc BannerController) GetUserBanner(c *gin.Context) {
	getUserBanner, err := banner.GetUserBanner(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Message},
		)
		return
	}

	content, err := bc.bannerService.GetUserBanner(*getUserBanner)

	if err != nil {

		if err.Status == 404 {
			c.String(http.StatusNotFound, "")
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Message},
		)
		return
	}

	c.JSON(http.StatusOK, content)
}

func (bc BannerController) CreateBanner(c *gin.Context) {

	modelBanner, err := bc.bannerValidator.ValidationCreateBanner(c)
	
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Message},
		)
		return
	}

	id, err := bc.bannerService.CreateBanner(*modelBanner)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Message},
		)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"banner_id": id})
}

func (bc BannerController) UpdateBanner(c *gin.Context) {

	updateBanner, err := bc.bannerValidator.ValidationUpdateBanner(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, 
			gin.H{"error": err.Message},
		)
		return
	}

	err = bc.bannerService.UpdateBanner(*updateBanner)

	if err != nil {
		if err.Status == 404 {
			c.String(http.StatusNotFound, "")
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Message},
		)
		return
	}

	c.String(http.StatusOK, "")
}

func (bc BannerController) DeleteBanner(c *gin.Context) {
	
	modelBanner := model.Banner{}
	goErr := c.BindUri(&modelBanner)

	if goErr != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "incorrect id"},
		)
		return
	}

	err := bc.bannerService.DeleteBanner(modelBanner)

	if err != nil {

		if err.Status == 404 {
			c.String(http.StatusNotFound, "")
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Message},
		)
		return
	}

	c.String(http.StatusNoContent, "")
}

func NewController(bs service.BannerService, bv validator.BannerValidator) *BannerController {
	return &BannerController{
		bannerService: bs,
		bannerValidator: bv,
	}
}