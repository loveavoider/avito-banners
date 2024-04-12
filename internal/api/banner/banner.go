package banner

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/converter/banner"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/service"
)

type BannerController struct {
	bannerService service.BannerService
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
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Message},
		)
		return
	}

	c.JSON(http.StatusOK, content)
}

func (bc BannerController) CreateBanner(c *gin.Context) {

	admin, status := bc.isAdmin(c)
	
	if !admin {
		c.String(status, "")
		return
	}

	modelBanner, err := banner.FromJsonToModel(c)

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

	admin, status := bc.isAdmin(c)

	if !admin {
		c.String(status, "")
		return
	}

	modelBanner, err := banner.FromJsonToModel(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, 
			gin.H{"error": err.Message},
		)
		return
	}

	err = bc.bannerService.UpdateBanner(*modelBanner)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Message},
		)
		return
	}

	c.String(http.StatusOK, "")
}

func (bc BannerController) DeleteBanner(c *gin.Context) {

	admin, status := bc.isAdmin(c)

	if !admin {
		c.String(status, "")
		return
	}
	
	id, goErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if goErr != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "incorrect id"},
		)
		return
	}

	modelBanner := model.Banner{ID: uint(id)}

	err := bc.bannerService.DeleteBanner(modelBanner)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Message},
		)
		return
	}

	c.String(http.StatusNoContent, "")
}

func (bc BannerController) isAdmin(c *gin.Context) (isAdmin bool, status int) {
	role, ok := c.Get("aud")

	if !ok {
		return false, http.StatusUnauthorized
	}

	log.Println("from isadmin", role)

	if role != "admin" {
		return false, http.StatusForbidden
	}

	return true, 0
}

func NewController(bs service.BannerService) *BannerController {
	return &BannerController{
		bannerService: bs,
	}
}