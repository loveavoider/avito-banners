package banner

import (
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/internal/repository/banner/converter"
	"github.com/loveavoider/avito-banners/internal/repository/banner/entity"
	"github.com/loveavoider/avito-banners/merror"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetBanners(getBanners model.GetBanners) (banners []model.BannerResponse, err *merror.MError) {
	entityBanners := make([]entity.Banner, 0)

	limit := -1
	if getBanners.Limit != 0 {
		limit = getBanners.Limit
	}

	res := r.db.Limit(limit).Offset(getBanners.Offset).Model(&entity.Banner{}).Preload("Tags").Find(&entityBanners)

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

	gormErr := r.db.Limit(limit).Offset(getBanners.Offset).Model(&entity.Tag{ID: getBanners.TagId}).Association("Banners").Find(&entityBanners)

	if gormErr != nil {
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

	res := r.db.Limit(limit).Offset(getBanners.Offset).Model(&entity.Banner{}).Preload("Tags").Find(&entityBanners, 
			&entity.Banner{FeatureId: getBanners.FeatureId})

	if res.Error != nil {
		return banners, &merror.MError{Message: "get banners by feature error"}
	}

	for _, banner := range entityBanners {
		banners = append(banners, converter.FromEntityToResponse(banner))
	}

	return
}

func (r *repository) GetUserBanner(tagId uint, featureId int) (content model.BannerResponse, err *merror.MError) {
	banner := entity.Banner{}

	subQuery := r.db.Select("banner_id").Where("tag_id = ?", tagId).Table("banner_tags")
	res := r.db.Where("ID IN (?) AND feature_id = ?", subQuery, featureId).First(&banner)

	if res.Error != nil {
		return content, &merror.MError{Message: "get banner error"}
	}

	content = converter.FromEntityToResponse(banner)

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

func (r *repository) UpdateBanner(bannerModel model.Banner) (err *merror.MError) {

	bannerEntitty := converter.FromModelToEntity(bannerModel)

	if len(*bannerEntitty.Tags) > 0 {
		err := r.updateTags(&bannerEntitty)

		if err != nil {
			return err
		}
	}

	res := r.db.Save(&bannerEntitty)

	if res.Error != nil {
		return &merror.MError{Message: "update banner error"}
	}

	return nil
}

func (r *repository) DeleteBanner(bannerModel model.Banner) (err *merror.MError) {
	bannerEntitty := converter.FromModelToEntity(bannerModel)

	res := r.db.Select(clause.Associations).Delete(&bannerEntitty)

	if res.Error != nil {
		return &merror.MError{Message: "delete banner error"}
	}

	return nil
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
