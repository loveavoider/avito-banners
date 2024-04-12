package app

import (
	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/api"
	"github.com/loveavoider/avito-banners/internal/repository"
	"github.com/loveavoider/avito-banners/internal/service"
	"github.com/loveavoider/avito-banners/internal/storage/database"

	bannerController "github.com/loveavoider/avito-banners/internal/api/banner"
	"github.com/loveavoider/avito-banners/internal/api/middleware"
	tokenController "github.com/loveavoider/avito-banners/internal/api/token"
	bannerRepository "github.com/loveavoider/avito-banners/internal/repository/banner"
	bannerService "github.com/loveavoider/avito-banners/internal/service/banner"
	tokenService "github.com/loveavoider/avito-banners/internal/service/token"
)

type serviceProvider struct {
	bannerRepository repository.BannerRepository
	bannerService service.BannerService
	bannerController api.BannerController

	tokenController api.TokenController
	tokenService service.TokenService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) BannerRepository() repository.BannerRepository {
	if s.bannerRepository == nil {
		s.bannerRepository = bannerRepository.NewRepository(database.NewPostgres())
	}

	return s.bannerRepository
}

func (s *serviceProvider) BannerService() service.BannerService {
	if s.bannerService == nil {
		s.bannerService = bannerService.NewService(
			s.BannerRepository(),
		)
	}

	return s.bannerService
}

func (s *serviceProvider) BannerController() api.BannerController {
	if s.bannerController == nil {
		s.bannerController = bannerController.NewController(
			s.BannerService(),
		)
	}
	
	return s.bannerController
}

func (s *serviceProvider) TokenService() service.TokenService {
	if s.tokenService == nil {
		s.tokenService = tokenService.NewTokenService()
	}

	return s.tokenService
}

func (s *serviceProvider) TokenController() api.TokenController {
	if s.tokenController == nil {
		s.tokenController = tokenController.NewTockenController(s.TokenService())
	}
	
	return s.tokenController
}

func (s *serviceProvider) TokenValidator() gin.HandlerFunc {
	return middleware.TokenValidator(s.TokenService())
}