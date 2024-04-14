package banner_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	bannerController "github.com/loveavoider/avito-banners/internal/api/banner"
	"github.com/loveavoider/avito-banners/internal/api/middleware"
	"github.com/loveavoider/avito-banners/internal/converter/banner"
	"github.com/loveavoider/avito-banners/internal/model"
	repoMock "github.com/loveavoider/avito-banners/internal/repository/mocks"
	bannerService "github.com/loveavoider/avito-banners/internal/service/banner"
	tokenService "github.com/loveavoider/avito-banners/internal/service/token"
	bannerValidator "github.com/loveavoider/avito-banners/internal/validator/banner"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetBanner(t *testing.T) {
	os.Setenv("JWT_KEY", "secret")
	router := gin.Default()	

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockBannerRepository(ctl)

	m := model.GetBanners{TagId: 4, FeatureId: 5, Role: "user"}

	text := "text"
	content := "content"
	url := "url"

	expRes := model.BannerResponse{
		ID: 1,
		TagIds: []uint{4, 5},
		FeatureId: 5,
		IsActive: true,
		Content: model.BannerContent{
			Title: &text,
			Text: &content,
			Url: &url,
		},
	}

	repo.EXPECT().GetUserBannerWithTags(m, true).Return(expRes, nil).Times(1)

	service := bannerService.NewService(repo)

	converter := banner.NewBannerConverter()
	validator := bannerValidator.NewBannerValidator(*converter, service)
	controller := bannerController.NewController(service, validator)

	bannerGroup := router.Group("/banner").Use(middleware.TokenValidator(tokenService.NewTokenService()))
	bannerGroup.GET("", controller.GetBanners)

	req := httptest.NewRequest(http.MethodGet, "/banner?tag_id=4&feature_id=5", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyIn0.EtFB0Mf0i-eXLIArptTSczbXwNiBszr-ijGAMTIDsm8")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	res := w.Result()

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	expected := `[{"banner_id":1,"tag_ids":[4,5],"feature_id":5,"content":{"title":"text","text":"content","url":"url"},"is_active":true,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`

	require.Equal(t, expected, string(data))
}