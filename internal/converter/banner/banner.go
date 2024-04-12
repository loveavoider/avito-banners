package banner

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/loveavoider/avito-banners/internal/model"
	"github.com/loveavoider/avito-banners/merror"
)

func FromJsonToModel(c *gin.Context) (*model.Banner, *merror.MError) {

	res := model.Banner{}
	err := c.BindJSON(&res)

	if err != nil {
		return nil, &merror.MError{Message: "invalid json"}
	}

	if c.Param("id") != "" {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)

		if err != nil {
			return nil, &merror.MError{Message: "incorrect id"}
		}

		if id != 0 {
			res.ID = uint(id)
		}
	}

	return &res, nil
}

func GetUserBanner(c *gin.Context) (*model.GetUserBanner, *merror.MError) {

	required := []string{"tag_id", "feature_id"}

	TagId, err := parseIntQueryField(c, "tag_id", required)
	
	if err != nil {
		return nil, &merror.MError{Message: err.Message}
	}

	FeatureId, err := parseIntQueryField(c, "feature_id", required)

	if err != nil {
		return nil, &merror.MError{Message: err.Message}
	}

	return &model.GetUserBanner{
		TagId: uint(TagId),
		FeatureId: FeatureId,
	}, nil
}

func GetBanners(c *gin.Context) (*model.GetBanners, *merror.MError) {

	required := make([]string, 0)
	
	tagId, err := parseIntQueryField(c, "tag_id", required)
	if err != nil {
		return nil, &merror.MError{Message: err.Message}
	}

	featureId, err := parseIntQueryField(c, "feature_id", required)
	if err != nil {
		return nil, &merror.MError{Message: err.Message}
	}

	limit, err := parseIntQueryField(c, "limit", required)
	if err != nil {
		return nil, &merror.MError{Message: err.Message}
	}

	offset, err := parseIntQueryField(c, "offset", required)
	if err != nil {
		return nil, &merror.MError{Message: err.Message}
	}

	return &model.GetBanners{
		TagId: uint(tagId),
		FeatureId: featureId,
		Limit: limit,
		Offset: offset,
	}, nil
}

func parseIntQueryField(c *gin.Context, field string, required []string) (int, *merror.MError) {

	fieldVal := c.Query(field)

	if fieldVal == "" {
		if inRequired(field, required) {
			return 0, &merror.MError{Message: fmt.Sprintf("missing require field %s", field)}
		}

		return 0, nil
	}

	res, err := strconv.Atoi(fieldVal)

	if err != nil {
		return 0, &merror.MError{Message: fmt.Sprintf("incorrect field %s", field)}
	}

	return res, nil
}

func inRequired(field string, required []string) bool {
	
	for _, str := range required {
		if str == field {
			return true
		}
	}

	return false
}