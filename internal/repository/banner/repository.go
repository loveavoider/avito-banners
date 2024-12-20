package banner

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/repository/banner/converter"
	"github.com/loveavoider/avito-banners/internal/repository/banner/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	NoFieldsToUpdate = errors.New("no fields to update")
	DbError          = errors.New("")
	BannersNotFound  = errors.New("")
)

type repository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewRepository(db *gorm.DB, cache *redis.Client) *repository {
	return &repository{
		db:    db,
		cache: cache,
	}
}

func (r *repository) GetBanners(getBanners model.GetBanners) (banners []model.BannerResponse, err error) {
	entityBanners := make([]entity.Banner, 0)

	limit := -1
	if getBanners.Limit != 0 {
		limit = getBanners.Limit
	}

	model := &entity.Banner{}

	if getBanners.Role == "user" {
		model.IsActive = true
	}

	res := r.db.Limit(limit).Offset(getBanners.Offset).Model(&entity.Banner{}).Preload("Tags").Find(&entityBanners, model)

	if len(entityBanners) == 0 || errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return banners, BannersNotFound
	}

	if res.Error != nil {
		return banners, DbError
	}

	for _, banner := range entityBanners {
		banners = append(banners, converter.FromEntityToResponse(banner))
	}

	return
}

func (r *repository) GetBannersByTag(getBanners model.GetBanners) (banners []model.BannerResponse, err error) {
	entityBanners := make([]entity.Banner, 0)

	limit := -1
	if getBanners.Limit != 0 {
		limit = getBanners.Limit
	}

	subQuery := r.db.Select("banner_id").Where("tag_id = ?", getBanners.TagId).Table("banner_tags")

	var res *gorm.DB

	// TODO сделать красиво limit offset
	if getBanners.Role == "admin" {
		res = r.db.
			Limit(limit).
			Offset(getBanners.Offset).
			Where("ID IN (?)", subQuery).
			Preload("Tags").
			Find(&entityBanners)
	} else {
		res = r.db.
			Limit(limit).
			Offset(getBanners.Offset).
			Where("ID IN (?) AND is_active = true", subQuery).
			Preload("Tags").Find(&entityBanners)
	}

	if len(entityBanners) == 0 || errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return banners, BannersNotFound
	}

	if res.Error != nil {
		return banners, DbError
	}

	for _, banner := range entityBanners {
		banners = append(banners, converter.FromEntityToResponse(banner))
	}

	return
}

func (r *repository) GetBannersByFeature(getBanners model.GetBanners) (banners []model.BannerResponse, err error) {
	entityBanners := make([]entity.Banner, 0)

	limit := -1
	if getBanners.Limit != 0 {
		limit = getBanners.Limit
	}

	model := &entity.Banner{FeatureId: getBanners.FeatureId}

	if getBanners.Role == "user" {
		model.IsActive = true
	}

	res := r.db.Limit(limit).Offset(getBanners.Offset).Model(&entity.Banner{}).Preload("Tags").Find(&entityBanners, model)

	if len(entityBanners) == 0 || errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return banners, BannersNotFound
	}

	if res.Error != nil {
		return banners, DbError
	}

	for _, banner := range entityBanners {
		banners = append(banners, converter.FromEntityToResponse(banner))
	}

	return
}

func (r *repository) GetUserBannerWithTags(getBanners model.GetBanners, useCache bool) (banner model.BannerResponse, err error) {
	entityBanner := entity.Banner{}
	ctx := context.Background()

	cacheKey := fmt.Sprintf("gUbWt_%d_%d", getBanners.FeatureId, getBanners.TagId)

	if useCache {
		val, redisErr := r.cache.Get(ctx, cacheKey).Result()

		if redisErr == nil {
			json.Unmarshal([]byte(val), &banner)
			return
		}
	}

	subQuery := r.db.Select("banner_id").Where("tag_id = ?", getBanners.TagId).Table("banner_tags")

	var res *gorm.DB

	if getBanners.Role == "admin" {
		res = r.db.Where("ID IN (?) AND feature_id = ?", subQuery, getBanners.FeatureId).Preload("Tags").First(&entityBanner)
	} else {
		res = r.db.Where("ID IN (?) AND feature_id = ? AND is_active = true", subQuery, getBanners.FeatureId).Preload("Tags").First(&entityBanner)
	}

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return banner, BannersNotFound
		}
		return banner, DbError
	}

	if useCache {
		banner = converter.FromEntityToResponse(entityBanner)

		stringBanner, _ := json.Marshal(banner)
		redisErr := r.cache.Set(ctx, cacheKey, stringBanner, 5*time.Minute).Err()

		if redisErr != nil {
			log.Println(redisErr.Error())
		}
	}

	return
}

func (r *repository) GetUserBanner(getUserBanner model.GetUserBanner, useCache bool) (content model.BannerContent, err error) {
	banner := entity.Banner{}
	ctx := context.Background()
	cacheKey := fmt.Sprintf("gUb_%d_%d", getUserBanner.FeatureId, getUserBanner.TagId)

	if useCache {
		val, redisErr := r.cache.Get(ctx, cacheKey).Result()

		if redisErr == nil {
			json.Unmarshal([]byte(val), &content)
			return
		}
	}

	subQuery := r.db.Select("banner_id").Where("tag_id = ?", getUserBanner.TagId).Table("banner_tags")

	var res *gorm.DB

	if getUserBanner.Role == "admin" {
		res = r.db.Where("ID IN (?) AND feature_id = ?", subQuery, getUserBanner.FeatureId).First(&banner)
	} else {
		res = r.db.Where("ID IN (?) AND feature_id = ? AND is_active = true", subQuery, getUserBanner.FeatureId).First(&banner)
	}

	if res.Error != nil {

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return content, BannersNotFound
		}

		return content, DbError
	}

	if useCache {
		content = converter.ConvertEntityContent(banner)

		stringContent, _ := json.Marshal(content)
		redisErr := r.cache.Set(ctx, cacheKey, stringContent, 5*time.Minute).Err()

		if redisErr != nil {
			log.Println(redisErr.Error())
		}
	}

	content.Title = banner.Content.Title
	content.Text = banner.Content.Text
	content.Url = banner.Content.Url

	return
}

func (r *repository) CreateBanner(bannerModel model.Banner) (id uint, err error) {
	bannerEntity := converter.FromModelToEntity(bannerModel)

	res := r.db.Create(&bannerEntity)

	if res.Error != nil {
		return id, DbError
	}

	return bannerEntity.ID, nil
}

func (r *repository) UpdateBanner(bannerModel model.UpdateBanner) (err error) {

	bannerEntity, selectFields := converter.BannerUpdateFromModelToEntity(bannerModel)

	if len(selectFields) == 0 {
		return NoFieldsToUpdate
	}

	if bannerEntity.Tags != nil && len(*bannerEntity.Tags) > 0 {
		err := r.db.Model(&bannerEntity).Association("Tags").Replace(bannerEntity.Tags)

		if err != nil {
			return DbError
		}
	}

	res := r.db.Model(&bannerEntity).Select(selectFields).Updates(bannerEntity)

	if res.RowsAffected == 0 {
		return BannersNotFound
	}

	if res.Error != nil {
		return DbError
	}

	return nil
}

func (r *repository) DeleteBanner(bannerModel model.Banner) (err error) {
	bannerEntity := converter.FromModelToEntity(bannerModel)

	res := r.db.Select(clause.Associations).Delete(&bannerEntity)

	if res.RowsAffected == 0 {
		return BannersNotFound
	}

	if res.Error != nil {
		return DbError
	}

	return nil
}

func (r *repository) CheckUnique(featureId int) (tags []uint, err error) {
	subQuery := r.db.Select("id").Where("feature_id = ?", featureId).Table("banners")
	res := r.db.Distinct("tag_id").Where("banner_id IN (?)", subQuery).Table("banner_tags").Find(&tags)

	if res.Error != nil {
		return tags, DbError
	}

	return tags, nil
}

func (r *repository) CheckUniqueByFeature(bannerId uint) (tags []uint, err error) {

	res := r.db.Select("tag_id").Where("banner_id = ?", bannerId).Table("banner_tags").Find(&tags)

	if res.Error != nil {
		return tags, DbError
	}

	return tags, nil
}

func (r *repository) CheckUniqueByTags(tagIds []uint, bannerId uint) (isUnique bool, err error) {
	var features []int

	banner := model.Banner{ID: bannerId}
	// найти фичу баннера
	res := r.db.First(&banner)

	if res.Error != nil {
		return false, DbError
	}

	// найти фичи по тегам
	subQuery := r.db.Distinct().Select("banner_id").Where("tag_id IN ", tagIds).Table("banner_tags")
	res = r.db.Distinct("feature_id").Where("banner_id IN (?)", subQuery).Table("banners").Find(&features)

	if res.Error != nil {
		return false, DbError
	}

	for _, feature := range features {
		if feature == banner.FeatureId {
			return false, nil
		}
	}

	return true, nil
}
