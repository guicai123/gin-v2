package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guicai123/gin-v2/pkg/app"
	"github.com/guicai123/gin-v2/pkg/e"
	"github.com/guicai123/gin-v2/pkg/util"
	"github.com/guicai123/gin-v2/service/member_service"
	"net/http"
	"time"
)

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
	articleService := member_service.Member{
		Username: username,
		Password: password,
		Token:    token,
		Created:  now.Format("2006-01-02 15:04:05"),
	}
	fmt.Sprintf("%T", articleService)
	if err := articleService.RegMember(); err != nil {
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
