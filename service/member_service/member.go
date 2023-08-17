package member_service

import (
	"fmt"
	"github.com/guicai123/gin-v2/models"
)

type Member struct {
	ID       int
	PageNum  int
	PageSize int
	Username string
	Password string
	Token    string
	Created  string
}

// 注册用户信息
func (a *Member) RegMember() error {
	memberInfo := map[string]interface{}{
		"user_name": a.Username,
		"password":  a.Password,
		"token":     a.Token,
		"created":   a.Created,
	}
	if err := models.AddMember(memberInfo); err != nil {
		return err
	}
	return nil
}

// 获取单个用户信息
func (a *Member) GetOne() ([]models.Member, error) {
	var (
		tags []models.Member
	)
	tags, err := models.GetMemberOne(a.getMaps())
	if err != nil {
		return nil, err
	}
	fmt.Println(tags)
	return tags, nil
}

// 获取全部用户信息
func (a *Member) GetAll() ([]models.Member, error) {
	var (
		tags []models.Member
	)
	tags, err := models.GetMemberAll(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (a *Member) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if a.Username != "" {
		maps["user_name"] = a.Username
	}
	if a.Password != "" {
		maps["password"] = a.Password
	}
	maps["is_deleted"] = 1
	return maps
}
