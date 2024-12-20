package banner

import (
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/repository"
)

type service struct {
	bannerRepository repository.BannerRepository
}

func NewService(bannerRepository repository.BannerRepository) *service {
	return &service{
		bannerRepository: bannerRepository,
	}
}

func (s *service) GetBanners(getBanners model.GetBanners) (banners []model.BannerResponse, err error) {

	if getBanners.TagId != 0 {

		if getBanners.FeatureId != 0 {
			useCache := getBanners.Role == "user"

			banner, err := s.bannerRepository.GetUserBannerWithTags(getBanners, useCache)

			if err != nil {
				return banners, err
			}

			banners = append(banners, banner)

			return banners, nil
		}

		banners, err = s.bannerRepository.GetBannersByTag(getBanners)

		if err != nil {
			return banners, err
		}

		return
	}

	if getBanners.FeatureId != 0 {
		banners, err = s.bannerRepository.GetBannersByFeature(getBanners)

		if err != nil {
			return banners, err
		}

		return
	}

	banners, err = s.bannerRepository.GetBanners(getBanners)

	if err != nil {
		return banners, err
	}

	return
}

func (s *service) GetUserBanner(getUserBanner model.GetUserBanner) (content model.BannerContent, err error) {
	useCache := !getUserBanner.UseLastRevision && getUserBanner.Role == "user"

	content, err = s.bannerRepository.GetUserBanner(getUserBanner, useCache)

	if err != nil {
		return content, err
	}

	return content, nil
}

func (s *service) CreateBanner(banner model.Banner) (id uint, err error) {
	return s.bannerRepository.CreateBanner(banner)
}

func (s *service) UpdateBanner(banner model.UpdateBanner) error {
	return s.bannerRepository.UpdateBanner(banner)
}

func (s *service) DeleteBanner(banner model.Banner) error {
	return s.bannerRepository.DeleteBanner(banner)
}

func (s *service) CheckUniqueByFeature(featureId int, bannerId uint) bool {
	tags, err := s.bannerRepository.CheckUniqueByFeature(bannerId)

	if err != nil {
		return false
	}

	return s.CheckUnique(featureId, tags)
}

func (s *service) CheckUniqueByTags(tagIds []uint, userId uint) bool {
	isUnique, err := s.bannerRepository.CheckUniqueByTags(tagIds, userId)

	if err != nil {
		return false
	}

	return isUnique
}

func (s *service) CheckUnique(featureId int, tagIds []uint) bool {
	tags, err := s.bannerRepository.CheckUnique(featureId)

	if err != nil {
		return false
	}

	store := make(map[uint]int)

	for _, tagId := range tags {
		store[tagId] += 1
	}

	for _, tagId := range tagIds {
		store[tagId] += 1
		if store[tagId] == 2 {
			return false
		}
	}

	return true
}
