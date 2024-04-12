package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/merror"
)

type BannerService interface {
	CreateBanner(model.Banner) (uint, *merror.MError)
	DeleteBanner(model.Banner) *merror.MError
	UpdateBanner(model.Banner) *merror.MError
	GetBanners(model.GetBanners) ([]model.BannerResponse, *merror.MError)
	GetUserBanner(model.GetUserBanner) (model.BannerContent, *merror.MError)
}

type TokenService interface {
	Generate(string) (string, *merror.MError)
	Validate(string) (*jwt.MapClaims, *merror.MError)
}