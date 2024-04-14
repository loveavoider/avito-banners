package repository

import (
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/merror"
)

type BannerRepository interface {
	CreateBanner(model.Banner) (uint, *merror.MError)
	DeleteBanner(model.Banner) *merror.MError
	UpdateBanner(model.UpdateBanner) *merror.MError
	GetBanners(model.GetBanners) ([]model.BannerResponse, *merror.MError)
	GetBannersByTag(model.GetBanners) ([]model.BannerResponse, *merror.MError)
	GetBannersByFeature(model.GetBanners) ([]model.BannerResponse, *merror.MError)
	GetUserBanner(model.GetUserBanner, bool) (model.BannerContent, *merror.MError)
	GetUserBannerWithTags(model.GetBanners, bool) (model.BannerResponse, *merror.MError)
	CheckUnique(int) ([]uint, *merror.MError)
}