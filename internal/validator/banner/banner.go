package banner

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/converter/banner"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/service"
)

type BannerValidator struct {
	bannerConverter banner.BannerConverter
	bannerService   service.BannerService
}

var (
	EmptyTagList       = errors.New("tag list should not be empty")
	EmptyTitle         = errors.New("title should not be empty")
	EmptyText          = errors.New("text should not be empty")
	EmptyUrl           = errors.New("url should not be empty")
	IncorrectTagList   = errors.New("incorrect tag list")
	IncorrectFeatureId = errors.New("incorrect feature id")
	IncorrectJson      = errors.New("incorrect json")
	UniqueTagFeature   = errors.New("combination tag + feature should be unique")
)

// TODO вынести общие функции в одну

func (bv *BannerValidator) ValidationCreateBanner(c *gin.Context) (*model.Banner, error) {

	b, err := bv.bannerConverter.FromJsonToBanner(c)

	if err != nil {
		return nil, IncorrectJson
	}

	if len(b.TagIds) == 0 {
		return nil, EmptyTagList
	}

	for _, tag := range b.TagIds {
		if tag == 0 {
			return nil, IncorrectTagList
		}
	}

	if b.FeatureId < 1 {
		return nil, IncorrectFeatureId
	}

	if b.Content.Title == nil || len(*b.Content.Title) == 0 {
		return nil, EmptyTitle
	}

	if b.Content.Text == nil || len(*b.Content.Text) == 0 {
		return nil, EmptyText
	}

	if b.Content.Url == nil || len(*b.Content.Url) == 0 {
		return nil, EmptyUrl
	}

	unique := bv.bannerService.CheckUnique(b.FeatureId, b.TagIds)

	if !unique {
		return nil, UniqueTagFeature
	}

	return b, nil
}

func (bv *BannerValidator) ValidationUpdateBanner(c *gin.Context) (*model.UpdateBanner, error) {
	updateBanner, err := bv.bannerConverter.FromJsonToUpdateBanner(c)

	if err != nil {
		return nil, IncorrectJson
	}

	if updateBanner.TagIds != nil && updateBanner.FeatureId != nil {
		unique := bv.bannerService.CheckUnique(*updateBanner.FeatureId, *updateBanner.TagIds)

		if !unique {
			return nil, UniqueTagFeature
		}
	}

	if updateBanner.TagIds != nil {

		if len(*updateBanner.TagIds) == 0 {
			return nil, EmptyTagList
		}

		for _, tag := range *updateBanner.TagIds {
			if tag == 0 {
				return nil, IncorrectTagList
			}
		}

		unique := bv.bannerService.CheckUniqueByTags(*updateBanner.TagIds, updateBanner.ID)
		if !unique {
			return nil, UniqueTagFeature
		}
	}

	if updateBanner.FeatureId != nil {
		if *updateBanner.FeatureId < 1 {
			return nil, IncorrectFeatureId
		}

		unique := bv.bannerService.CheckUniqueByFeature(*updateBanner.FeatureId, updateBanner.ID)

		if !unique {
			return nil, UniqueTagFeature
		}
	}

	if updateBanner.Content != nil {
		if updateBanner.Content.Title != nil && len(*updateBanner.Content.Title) == 0 {
			return nil, EmptyTitle
		}

		if updateBanner.Content.Text != nil && len(*updateBanner.Content.Text) == 0 {
			return nil, EmptyText
		}

		if updateBanner.Content.Url != nil && len(*updateBanner.Content.Url) == 0 {
			return nil, EmptyUrl
		}
	}

	return updateBanner, nil
}

func NewBannerValidator(bannerConverter banner.BannerConverter, bannerService service.BannerService) *BannerValidator {
	return &BannerValidator{
		bannerConverter: bannerConverter,
		bannerService:   bannerService,
	}
}
