package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/merror"
)

type BannerService interface {
	CreateBanner(model.Banner) (uint, *merror.MError)
	DeleteBanner(model.Banner) *merror.MError
	UpdateBanner(model.UpdateBanner) *merror.MError
	GetBanners(model.GetBanners) ([]model.BannerResponse, *merror.MError)
	GetUserBanner(model.GetUserBanner) (model.BannerContent, *merror.MError)
	CheckUnique(int, []uint) bool
	CheckUniqueByFeature(int, uint) bool
	CheckUniqueByTags([]uint, uint) bool
}

type TokenService interface {
	Generate(string) (string, *merror.MError)
	Validate(string) (*jwt.MapClaims, *merror.MError)
}