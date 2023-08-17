package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guicai123/gin-v2/models"
	"github.com/guicai123/gin-v2/pkg/app"
	"github.com/guicai123/gin-v2/pkg/e"
	"github.com/guicai123/gin-v2/pkg/util"
	"github.com/guicai123/gin-v2/service/member_service"
	"net/http"
	"time"
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

// 用户注册信息
func DoReg(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	username := c.PostForm("username")
	password := c.PostForm("password")

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	now := time.Now()
	memberInfo := member_service.Member{
		Username: username,
		Password: util.EncodeMD5(password),
		Token:    token,
		Created:  now.Format("2006-01-02 15:04:05"),
	}

	//判断用户是否存在
	maps := make(map[string]interface{})
	maps["user_name"] = memberInfo.Username

	var (
		tags []models.Member
	)

	tags, err = models.GetMemberOne(maps)
	if err != nil {

	}

	if len(tags) == 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	fmt.Println(maps)
	fmt.Println(tags)

	//memberInfo2 := map[string]interface{}{
	//	"user_name": memberInfo.Username,
	//	"password":  memberInfo.Password,
	//	"token":     memberInfo.Token,
	//	"created":   memberInfo.Created,
	//}
	//if err := models.AddMember(memberInfo2); err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
	//	return
	//}

	fmt.Sprintf("%T", memberInfo)
	if err := memberInfo.RegMember(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 用户登陆
func DoLogin(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	username := c.PostForm("username")
	Password := c.PostForm("password")
	articleService := member_service.Member{
		Username: username,
		Password: Password,
	}
	member, err := articleService.GetOne()
	if err != nil || len(member) == 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DOLOGIN_FAIL, nil)
		return
	}
	memberInfo := make(map[string]interface{})
	memberInfo["id"] = member[0].ID
	memberInfo["username"] = member[0].UserName
	memberInfo["password"] = member[0].Password
	memberInfo["phone"] = member[0].Phone
	appG.Response(http.StatusOK, e.SUCCESS, memberInfo)
}
