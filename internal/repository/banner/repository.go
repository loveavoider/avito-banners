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
	"github.com/loveavoider/avito-banners/merror"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
	cache *redis.Client
}

func NewRepository(db *gorm.DB, cache *redis.Client) *repository {
	return &repository{
		db: db,
		cache: cache,
	}
}

func (r *repository) GetBanners(getBanners model.GetBanners) (banners []model.BannerResponse, err *merror.MError) {
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
		return banners, &merror.MError{Message: "", Status: 404}
	}

	if res.Error != nil {
		return banners, &merror.MError{Message: "get all banners error"}
	}

	for _, banner := range entityBanners {
		banners = append(banners, converter.FromEntityToResponse(banner))
	}

	return
}

func (r *repository) GetBannersByTag(getBanners model.GetBanners) (banners []model.BannerResponse, err *merror.MError) {
	entityBanners := make([]entity.Banner, 0)

	limit := -1
	if getBanners.Limit != 0 {
		limit = getBanners.Limit
	}

	subQuery := r.db.Select("banner_id").Where("tag_id = ?", getBanners.TagId).Table("banner_tags")

	var res *gorm.DB

	if getBanners.Role == "admin" { 
		res = r.db.Limit(limit).Offset(getBanners.Offset).Where("ID IN (?)", subQuery).Preload("Tags").Find(&entityBanners)
	} else {
		res = r.db.Limit(limit).Offset(getBanners.Offset).Where("ID IN (?) AND is_active = true", subQuery).Preload("Tags").Find(&entityBanners)
	}

	if len(entityBanners) == 0 || errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return banners, &merror.MError{Message: "", Status: 404}
	}

	if res.Error != nil {
		return banners, &merror.MError{Message: "get banners by tag error"}
	}

	for _, banner := range entityBanners {
		banners = append(banners, converter.FromEntityToResponse(banner))
	}

	return
}

func (r *repository) GetBannersByFeature(getBanners model.GetBanners) (banners []model.BannerResponse, err *merror.MError) {
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
		return banners, &merror.MError{Message: "", Status: 404}
	}

	if res.Error != nil {
		return banners, &merror.MError{Message: "get banners by feature error"}
	}

	for _, banner := range entityBanners {
		banners = append(banners, converter.FromEntityToResponse(banner))
	}

	return
}

func (r *repository) GetUserBannerWithTags(getBanners model.GetBanners, useCache bool) (banner model.BannerResponse, err *merror.MError)   {
	entityBanner := entity.Banner{}
	ctx := context.Background()

	cacheKey := fmt.Sprintf("gUbWt_%d_%d", getBanners.FeatureId, getBanners.TagId)

	if useCache {
		val, redisErr := r.cache.Get(ctx, cacheKey).Result()
		log.Println("from cache", val)
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
			return banner, &merror.MError{Message: "", Status: 404}
		}
		return banner, &merror.MError{Message: "get banner error"}
	}

	if useCache {
		banner = converter.FromEntityToResponse(entityBanner)

		stringBanner, _ := json.Marshal(banner)
		redisErr := r.cache.Set(ctx, cacheKey, stringBanner, 5 * time.Minute).Err()

		if redisErr != nil {
			log.Println(redisErr.Error())
		}
	}

	return
}

func (r *repository) GetUserBanner(getUserBanner model.GetUserBanner, useCache bool) (content model.BannerContent, err *merror.MError) {
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
			return content, &merror.MError{Message: "", Status: 404}
		}

		return content, &merror.MError{Message: "get banner error"}
	}

	if useCache {
		content = converter.ConvertEntityContent(banner)

		stringContent, _ := json.Marshal(content)
		redisErr := r.cache.Set(ctx, cacheKey, stringContent, 5 * time.Minute).Err()

		if redisErr != nil {
			log.Println(redisErr.Error())
		}
	}
	
	return
}

func (r *repository) CreateBanner(bannerModel model.Banner) (id uint, err *merror.MError) {
	bannerEntitty := converter.FromModelToEntity(bannerModel)

	res := r.db.Create(&bannerEntitty)

	if res.Error != nil {
		return id, &merror.MError{Message: "create error"}
	}

	return bannerEntitty.ID, nil
}

func (r *repository) UpdateBanner(bannerModel model.UpdateBanner) (err *merror.MError) {

	bannerEntitty, selectFields := converter.BannerUpdateFromModelToEntity(bannerModel)
	
	if len(selectFields) == 0 {
		return &merror.MError{Message: "Нет полей на обновление"}
	}
	
	if bannerEntitty.Tags != nil && len(*bannerEntitty.Tags) > 0 {
		err := r.updateTags(&bannerEntitty)

		if err != nil {
			return err
		}
	}

	res := r.db.Model(&bannerEntitty).Select(selectFields).Updates(bannerEntitty)

	if res.RowsAffected == 0 {
		return &merror.MError{Message: "", Status: 404}
	}

	if res.Error != nil {
		return &merror.MError{Message: "update banner error"}
	}

	return nil
}

func (r *repository) DeleteBanner(bannerModel model.Banner) (err *merror.MError) {
	bannerEntitty := converter.FromModelToEntity(bannerModel)

	res := r.db.Select(clause.Associations).Delete(&bannerEntitty)

	if res.RowsAffected == 0 {
		return &merror.MError{Message: "", Status: 404}
	}

	if res.Error != nil {
		return &merror.MError{Message: "delete banner error"}
	}

	return nil
}

func (r *repository) CheckUnique(featureId int) (tags []uint, err *merror.MError) {
	subQuery := r.db.Select("id").Where("feature_id = ?", featureId).Table("banners")
	res := r.db.Distinct("tag_id").Where("banner_id IN (?)", subQuery).Table("banner_tags").Find(&tags)

	if res.Error != nil {
		return tags, &merror.MError{Message: "check unique error"}
	}

	return tags, nil
}

func (r *repository) updateTags(bannerEntity *entity.Banner) (err *merror.MError) {
	var tags []entity.Tag
	gormErr := r.db.Model(&bannerEntity).Association("Tags").Find(&tags)

	if gormErr != nil {
		return &merror.MError{Message: "update tags error"}
	}

	tagsToDel := make([]entity.Tag, 0)

	for _, tag := range tags {
		if !r.haveTag(tag, *bannerEntity.Tags) {
			tagsToDel = append(tagsToDel, tag)
		}
	}

	if len(tagsToDel) > 0 {
		r.db.Model(&bannerEntity).Association("Tags").Delete(tagsToDel)
	}

	return nil
}

func (r *repository) haveTag(needTag entity.Tag, tags []entity.Tag) bool {

	for _, tag := range tags {
		if tag.ID == needTag.ID {
			return true
		}
	}

	return false
}
