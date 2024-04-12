package converter

import (
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/repository/banner/entity"
)

func FromEntityToModel (banner entity.Banner) model.Banner {
	return model.Banner{}
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
		ID: banner.ID,
		TagIds: tags,
		FeatureId: banner.FeatureId,
		Content: ConvertEntityContent(banner.Content.Title, banner.Content.Text, banner.Content.Url),
		IsActive: banner.IsActive,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}
}

func ConvertEntityContent(title string, text string, url string) model.BannerContent {
	return model.BannerContent{
		Title: title,
		Text: text,
		Url: url,
	}
}

func convertContent(content model.BannerContent) *entity.BannerContent {
	
	return &entity.BannerContent{
		Title: content.Title,
		Text: content.Text,
		Url: content.Url,
	}
}