package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/repository"
)

var (
	ErrNoFieldsToUpdate = repository.NoFieldsToUpdate
	DbError             = repository.DbError
	BannersNotFound     = repository.BannersNotFound
)

type BannerService interface {
	CreateBanner(model.Banner) (uint, error)
	DeleteBanner(model.Banner) error
	UpdateBanner(model.UpdateBanner) error
	GetBanners(model.GetBanners) ([]model.BannerResponse, error)
	GetUserBanner(model.GetUserBanner) (model.BannerContent, error)
	CheckUnique(int, []uint) bool
	CheckUniqueByFeature(int, uint) bool
	CheckUniqueByTags([]uint, uint) bool
}

type TokenService interface {
	Generate(string) (string, error)
	Validate(string) (*jwt.MapClaims, error)
}
