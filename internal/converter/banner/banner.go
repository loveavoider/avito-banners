package banner

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/model"
)

type BannerConverter struct {
}

func (bc *BannerConverter) FromJsonToBanner(c *gin.Context) (*model.Banner, error) {

	res := model.Banner{
		IsActive: true,
	}

	err := c.BindJSON(&res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (bc *BannerConverter) FromJsonToUpdateBanner(c *gin.Context) (*model.UpdateBanner, error) {

	res := model.UpdateBanner{}

	err := c.BindUri(&res)

	if err != nil {
		return nil, err
	}

	err = c.BindJSON(&res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func GetUserBanner(c *gin.Context) (*model.GetUserBanner, error) {

	role, _ := c.Get("aud")

	// TODO throw 401

	getUserBanner := &model.GetUserBanner{Role: role.(string)}
	err := c.BindQuery(&getUserBanner)

	// TODO перенести в валидатор

	if err != nil {
		return nil, err
	}

	if getUserBanner.FeatureId == 0 || getUserBanner.TagId == 0 {
		return nil, errors.New("incorrect id or tag id")
	}

	return getUserBanner, nil
}

func GetBanners(c *gin.Context) (*model.GetBanners, error) {

	role, _ := c.Get("aud")

	getBanners := &model.GetBanners{Role: role.(string)}
	err := c.BindQuery(&getBanners)

	if err != nil {
		return nil, err
	}

	return getBanners, nil
}

func NewBannerConverter() *BannerConverter {
	return &BannerConverter{}
}
