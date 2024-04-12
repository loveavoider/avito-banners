package banner

import (
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/repository"
	"github.com/loveavoider/avito-banners/merror"
)

type service struct {
	bannerRepository repository.BannerRepository
}

func NewService(bannerRepository repository.BannerRepository) *service {
	return &service{
		bannerRepository: bannerRepository,
	}
}

func (s *service) GetBanners(getBanners model.GetBanners) (banners []model.BannerResponse, err *merror.MError) {

	if getBanners.TagId != 0 {

		if getBanners.FeatureId != 0 {
			banner, err := s.bannerRepository.GetUserBanner(getBanners.TagId, getBanners.FeatureId)

			if err != nil {
				return banners, &merror.MError{Message: err.Message}
			}

			banners = append(banners, banner)

			return banners, nil
		}

		banners, err = s.bannerRepository.GetBannersByTag(getBanners)

		if err != nil {
			return banners, &merror.MError{Message: err.Message}
		}

		return
	}

	if getBanners.FeatureId != 0 {
		banners, err = s.bannerRepository.GetBannersByFeature(getBanners)

		if err != nil {
			return banners, &merror.MError{Message: err.Message}
		}

		return
	}

	banners, err = s.bannerRepository.GetBanners(getBanners)

	if err != nil {
		return banners, &merror.MError{Message: err.Message}
	}

	return
}

func (s *service) GetUserBanner(getUserBanner model.GetUserBanner) (content model.BannerContent, err *merror.MError) {
	banner, err := s.bannerRepository.GetUserBanner(getUserBanner.TagId, getUserBanner.FeatureId)
	content = banner.Content

	return
}

func (s *service) CreateBanner(banner model.Banner) (id uint, err *merror.MError) {
	id, err = s.bannerRepository.CreateBanner(banner)

	return
}

func (s *service) UpdateBanner(banner model.Banner) (err *merror.MError) {
	err = s.bannerRepository.UpdateBanner(banner)

	return
}

func (s *service) DeleteBanner(banner model.Banner) (err *merror.MError) {
	err = s.bannerRepository.DeleteBanner(banner)

	return
}