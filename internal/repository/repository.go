package repository

import (
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/repository/banner"
)

var (
	NoFieldsToUpdate = banner.NoFieldsToUpdate
	DbError          = banner.DbError
	BannersNotFound  = banner.BannersNotFound
)

type BannerRepository interface {
	CreateBanner(model.Banner) (uint, error)
	DeleteBanner(model.Banner) error
	UpdateBanner(model.UpdateBanner) error
	GetBanners(model.GetBanners) ([]model.BannerResponse, error)
	GetBannersByTag(model.GetBanners) ([]model.BannerResponse, error)
	GetBannersByFeature(model.GetBanners) ([]model.BannerResponse, error)
	GetUserBanner(model.GetUserBanner, bool) (model.BannerContent, error)
	GetUserBannerWithTags(model.GetBanners, bool) (model.BannerResponse, error)
	CheckUnique(int) ([]uint, error)
	CheckUniqueByFeature(uint) ([]uint, error)
	CheckUniqueByTags([]uint, uint) (bool, error)
}
