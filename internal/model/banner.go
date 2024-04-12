package model

import "time"

type BannerContent struct {
	Title string `json:"title"`
	Text string `json:"text"`
	Url string `json:"url"`
}

type Banner struct {
	ID uint
	TagIds []uint `json:"tag_ids"`
	FeatureId int `json:"feature_id"`
	Content BannerContent `json:"content"`
	IsActive bool `json:"is_active"`
}

type BannerResponse struct {
	ID uint `json:"banner_id"`
	TagIds []uint `json:"tag_ids"`
	FeatureId int `json:"feature_id"`
	Content BannerContent `json:"content"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserBanner struct {
	TagId uint
	FeatureId int
}

type GetBanners struct {
	TagId uint
	FeatureId int
	Limit int
	Offset int
}