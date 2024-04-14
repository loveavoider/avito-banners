package banner

import (
	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/converter/banner"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/service"
	"github.com/loveavoider/avito-banners/merror"
)

type BannerValidator struct {
	bannerConverter banner.BannerConverter
	bannerService service.BannerService
}

func (bv *BannerValidator) ValidationCreateBanner(c *gin.Context) (*model.Banner, *merror.MError) {

	banner, err := bv.bannerConverter.FromJsonToBanner(c)

	if err != nil {
		return nil, &merror.MError{Message: err.Message}
	}

	if len(banner.TagIds) == 0 {
		return nil, &merror.MError{Message: "Список тегов не должен быть пустым"}
	}

	for _, tag := range banner.TagIds {
		if tag == 0 {
			return nil, &merror.MError{Message: "Некорректный список тегов"}
		}
	}

	if banner.FeatureId < 1 {
		return nil, &merror.MError{Message: "Некорректный id фичи"}
	}

	if banner.Content.Title == nil || len(*banner.Content.Title) == 0 {
		return nil, &merror.MError{Message: "Заголовок не должен быть пустым"}
	}

	if banner.Content.Text == nil || len(*banner.Content.Text) == 0 {
		return nil, &merror.MError{Message: "Текст не должен быть пустым"}
	}

	if banner.Content.Url == nil || len(*banner.Content.Url) == 0 {
		return nil, &merror.MError{Message: "Ссылка не должна быть пустой"}
	}

	unique := bv.bannerService.CheckUnique(banner.FeatureId, banner.TagIds)
	
	if !unique {
		return nil, &merror.MError{Message: "Набор feature_id + tag_id должен быть уникальным"}
	}

	return banner, nil
}

func (bv *BannerValidator) ValidationUpdateBanner(c *gin.Context) (*model.UpdateBanner, *merror.MError) {
	updateBanner, err := bv.bannerConverter.FromJsonToUpdateBanner(c)

	if err != nil {
		return nil, &merror.MError{Message: err.Message}
	}

	if updateBanner.TagIds != nil && updateBanner.FeatureId != nil {
		unique := bv.bannerService.CheckUnique(*updateBanner.FeatureId, *updateBanner.TagIds)
	
		if !unique {
			return nil, &merror.MError{Message: "Набор feature_id + tag_id должен быть уникальным"}
		}
	}

	if updateBanner.TagIds != nil {

		if len(*updateBanner.TagIds) == 0 {
			return nil, &merror.MError{Message: "Список тегов не должен быть пустым"}
		}

		for _, tag := range *updateBanner.TagIds {
			if tag == 0 {
				return nil, &merror.MError{Message: "Некорректный список тегов"}
			}
		}

		unique := bv.bannerService.CheckUniqueByTags(*updateBanner.TagIds, updateBanner.ID)
		if !unique {
			return nil, &merror.MError{Message: "Набор feature_id + tag_id должен быть уникальным"}
		}
	}

	if updateBanner.FeatureId != nil { 
		if *updateBanner.FeatureId < 1 {
			return nil, &merror.MError{Message: "Некорректный id фичи"}
		}

		unique := bv.bannerService.CheckUniqueByFeature(*updateBanner.FeatureId, updateBanner.ID)

		if !unique {
			return nil, &merror.MError{Message: "Набор feature_id + tag_id должен быть уникальным"}
		}
	}

	if updateBanner.Content != nil {
		if updateBanner.Content.Title != nil && len(*updateBanner.Content.Title) == 0 {
			return nil, &merror.MError{Message: "Заголовок не должен быть пустым"}
		}
	
		if updateBanner.Content.Text != nil && len(*updateBanner.Content.Text) == 0 {
			return nil, &merror.MError{Message: "Текст не должен быть пустым"}
		}
	
		if updateBanner.Content.Url != nil && len(*updateBanner.Content.Url) == 0 {
			return nil, &merror.MError{Message: "Ссылка не должна быть пустой"}
		}
	}

	return updateBanner, nil
}

func NewBannerValidator(bannerConverter banner.BannerConverter, bannerService service.BannerService) *BannerValidator {
	return &BannerValidator{
		bannerConverter: bannerConverter,
		bannerService: bannerService,
	}
}