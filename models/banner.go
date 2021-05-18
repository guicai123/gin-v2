package models

import "github.com/jinzhu/gorm"

type Banner struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`

}

//获取单个用户信息
func GetBannerOne(maps interface{}) ([]Banner, error) {
	var (
		banners []Banner
		err  error
	)
	err = db.Where(maps).Limit(5).Order("id Desc").Find(&banners).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return banners, nil
}
