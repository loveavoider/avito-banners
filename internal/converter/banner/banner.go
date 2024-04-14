package banner

import (
	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/merror"
)

type BannerConverter struct {
}

func (bc *BannerConverter) FromJsonToBanner(c *gin.Context) (*model.Banner, *merror.MError) {

	res := model.Banner{
		IsActive: true,
	}

	err := c.BindJSON(&res)

	if err != nil {
		return nil, &merror.MError{Message: "invalid json"}
	}

	return &res, nil
}

func (bc *BannerConverter) FromJsonToUpdateBanner(c *gin.Context) (*model.UpdateBanner, *merror.MError) {

	res := model.UpdateBanner{}

	err := c.BindUri(&res)

	if err != nil {
		return nil, &merror.MError{Message: "incorrect id"}
	}

	err = c.BindJSON(&res)

	if err != nil {
		return nil, &merror.MError{Message: "invalid json"}
	}

	return &res, nil
}

func GetUserBanner(c *gin.Context) (*model.GetUserBanner, *merror.MError) {

	role, _ := c.Get("aud")

	getUserBanner := &model.GetUserBanner{Role: role.(string)}
	ginErr := c.BindQuery(&getUserBanner)

	if ginErr != nil {
		return nil, &merror.MError{Message: ""}
	}

	if getUserBanner.FeatureId == 0 || getUserBanner.TagId == 0 {
		return nil, &merror.MError{Message: "неверно указан id"}
	}
	

	return getUserBanner, nil
}

func GetBanners(c *gin.Context) (*model.GetBanners, *merror.MError) {

	role, _ := c.Get("aud")

	getBanners := &model.GetBanners{Role: role.(string)}
	ginErr := c.BindQuery(&getBanners)

	if ginErr != nil {
		return nil, &merror.MError{Message: ""}
	}

	return getBanners, nil
}

func NewBannerConverter() *BannerConverter {
	return &BannerConverter{}
}