package index_service

import (
	"github.com/guicai123/gin-v2/models"
)

type IndexInfo struct {

}

func (a *IndexInfo) GetBannerOne() ([]models.Banner, error) {
	var (
		banners []models.Banner
	)
	maps:=make(map[string]interface{})
	maps["is_deleted"]="1"
	maps["cateId"]="16"
	banners, err := models.GetBannerOne(maps)
	if err != nil {
		return nil, err
	}
	return banners, nil
}

