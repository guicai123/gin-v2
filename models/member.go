package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Member struct {
	Model

	UserName string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Token    string `json:"token"`
	Created  string `json:"created"`
}

// 插入数据
func AddMember(data map[string]interface{}) error {
	member := Member{
		UserName:   data["username"].(string),
		Password:   data["password"].(string),
		Token:      data["token"].(string),
		Created:    data["created"].(string),
	}
	fmt.Println(data)
	if err := db.Create(&member).Error; err != nil {
		return err
	}
	return nil
}

//获取单个用户信息
func GetMemberOne(maps interface{}) ([]Member, error) {
	var (
		members []Member
		err  error
	)
	err = db.Where(maps).Find(&members).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return members, nil
}

//全部用户信息
func GetMemberAll(pageNum int, pageSize int, maps interface{}) ([]Member, error) {
	var (
		MemberAll []Member
		err  error
	)
	err = db.Select("id,name").Where(maps).Offset(pageNum).Limit(pageSize).Find(&MemberAll).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return MemberAll, nil
}




