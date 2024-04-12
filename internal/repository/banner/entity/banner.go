package entity

import (
	"time"
)

type Banner struct {
	ID uint `gorm:"primaryKey"`
	Tags *[]Tag `gorm:"many2many:banner_tags"`
	FeatureId int `gorm:"feature_id"`
	Content BannerContent `gorm:"embedded"`
	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tag struct {
	ID uint `gorm:"primaryKey"`
	Banners *[]Banner `gorm:"many2many:banner_tags"`
}

type BannerContent struct {
	Title string
	Text string
	Url string
}