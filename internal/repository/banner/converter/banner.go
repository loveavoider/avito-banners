package converter

import (
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/repository/banner/entity"
	"reflect"
)

func BannerUpdateFromModelToEntity(updateBanner model.UpdateBanner) (entity.Banner, []string) {
	res := entity.Banner{
		ID: updateBanner.ID,
	}

	selectFields := make([]string, 0, reflect.ValueOf(updateBanner).NumField())

	if updateBanner.TagIds != nil {
		tags := make([]entity.Tag, len(*updateBanner.TagIds))

		for i, tag := range *updateBanner.TagIds {
			tags[i].ID = tag
		}

		res.Tags = &tags
	}

	if updateBanner.Content != nil {
		res.Content = *convertContent(*updateBanner.Content)

		if res.Content.Title != nil {
			selectFields = append(selectFields, "title")
		}

		if res.Content.Text != nil {
			selectFields = append(selectFields, "text")
		}

		if res.Content.Url != nil {
			selectFields = append(selectFields, "url")
		}
	}

	if updateBanner.FeatureId != nil {
		res.FeatureId = *updateBanner.FeatureId
		selectFields = append(selectFields, "feature_id")
	}

	if updateBanner.IsActive != nil {
		res.IsActive = *updateBanner.IsActive
		selectFields = append(selectFields, "is_active")
	}

	return res, selectFields
}

func FromModelToEntity(banner model.Banner) entity.Banner {
	res := entity.Banner{}

	tags := make([]entity.Tag, len(banner.TagIds))

	if banner.ID != 0 {
		res.ID = banner.ID
	}

	res.Content = *convertContent(banner.Content)

	for i, tag := range banner.TagIds {
		tags[i].ID = tag
	}

	res.FeatureId = banner.FeatureId
	res.Tags = &tags
	res.IsActive = banner.IsActive

	return res
}

func FromEntityToResponse(banner entity.Banner) model.BannerResponse {
	var tags []uint

	if banner.Tags != nil {

		tags = make([]uint, len(*banner.Tags))

		for i, tag := range *banner.Tags {
			tags[i] = tag.ID
		}
	}

	return model.BannerResponse{
		ID:        banner.ID,
		TagIds:    tags,
		FeatureId: banner.FeatureId,
		Content:   ConvertEntityContent(banner),
		IsActive:  banner.IsActive,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}
}

func ConvertEntityContent(bannerEntity entity.Banner) model.BannerContent {
	return model.BannerContent{
		Title: bannerEntity.Content.Title,
		Text:  bannerEntity.Content.Text,
		Url:   bannerEntity.Content.Url,
	}
}

func convertContent(content model.BannerContent) *entity.BannerContent {
	return &entity.BannerContent{
		Title: content.Title,
		Text:  content.Text,
		Url:   content.Url,
	}
}
