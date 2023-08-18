package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type GoodsCates struct {
	Model
	Name string `json:"name"`
	//CreatedBy  string `json:"created_by"`
	//ModifiedBy string `json:"modified_by"`
	//State      int    `json:"state"`
}

// 获取全部信息
func GetGoodsCate(pageNum int, pageSize int, maps interface{}) ([]GoodsCates, error) {
	var (
		goodscate []GoodsCates
		err       error
	)
	fmt.Println(1)
	err = db.Select("id,name").Where(maps).Offset(pageNum).Limit(pageSize).Find(&goodscate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return goodscate, nil
}

// 统计平台数据
func GetGoodsCateTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&GoodsCates{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// DeleteTag delete a tag
func DeleteGoodsCate(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

// =========================2023/08/18============
func GetGoodsInfo() ([]GoodsCates, error) {
	var (
		GoodsCates []GoodsCates
		err        error
	)

	fmt.Println(1)
	//根据主键获取第一个条信息
	//err = db.First(&GoodsCates).Error

	//根据主键查询最后一条记录
	//err = db.Last(&GoodsCates).Error

	//随机获取一条记录
	//err = db.Take(&GoodsCates).Error
	//err = db.Order("RAND()").First(&GoodsCates).Error

	//查询所有的记录这里users定义的是数组
	//err = db.Find(&GoodsCates).Error
	//err = db.Find(&GoodsCates, 266).Error

	//where 条件进行查询
	//err = db.Where("name=?", "小家电").First(&GoodsCates).Error

	//err = db.Where("name=? and id=?", "小家电", "306").Find(&GoodsCates).Error

	//<>
	//err = db.Where("id <>? and name=?", "304", "小家电").Find(&GoodsCates).Error

	//IN
	//err = db.Where("id in (?) and name=?", []int{301, 302, 304, 305}, "小家电").Find(&GoodsCates).Error

	//NOTIN
	//err = db.Where("id not in (?) and name=?", []int{301, 302, 304, 305, 309}, "小家电").Find(&GoodsCates).Error

	//LIKE
	//err = db.Debug().Where("id not in (?) and name like ?", []int{301, 302, 304, 305, 309}, "%小家电%").Find(&GoodsCates).Error

	//AND
	//err = db.Where("id in (?) and name=?", []int{301, 302, 304, 305}, "小家电").Find(&GoodsCates).Error

	//运算符 >= <= !=
	// 获取当前时间
	now := time.Now()

	// 计算昨天的时间
	yesterday := now.AddDate(0, 0, -1)
	//err = db.Where("id >=? and add_time >=?", "320", yesterday).Find(&GoodsCates).Error

	//BETWEEN
	err = db.Where(" add_time between ? AND ?", yesterday, now).Find(&GoodsCates).Error
	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)

	//currtime := time.Now()

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return GoodsCates, nil

}
