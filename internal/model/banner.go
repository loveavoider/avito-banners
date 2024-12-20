package model

import "time"

type BannerContent struct {
	Title *string `json:"title"`
	Text  *string `json:"text"`
	Url   *string `json:"url"`
}

type Banner struct {
	ID        uint          `uri:"id"`
	TagIds    []uint        `json:"tag_ids"`
	FeatureId int           `json:"feature_id"`
	Content   BannerContent `json:"content"`
	IsActive  bool          `json:"is_active"`
}

type UpdateBanner struct {
	ID        uint           `uri:"id" binding:"required"`
	TagIds    *[]uint        `json:"tag_ids"`
	FeatureId *int           `json:"feature_id"`
	Content   *BannerContent `json:"content"`
	IsActive  *bool          `json:"is_active"`
}

// TODO Сделать сущностью баннер а не общим для всех слоев

type BannerResponse struct {
	ID        uint          `json:"banner_id"`
	TagIds    []uint        `json:"tag_ids"`
	FeatureId int           `json:"feature_id"`
	Content   BannerContent `json:"content"`
	IsActive  bool          `json:"is_active"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type GetUserBanner struct {
	TagId           uint `form:"tag_id" binding:"required"`
	FeatureId       int  `form:"feature_id" binding:"required"`
	UseLastRevision bool `form:"use_last_revision"`
	Role            string
}

type GetBanners struct {
	TagId     uint `form:"tag_id"`
	FeatureId int  `form:"feature_id"`
	Limit     int  `form:"limit"`
	Offset    int  `form:"offset"`
	Role      string
}
