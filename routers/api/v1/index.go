package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guicai123/gin-v2/pkg/app"
	"github.com/guicai123/gin-v2/pkg/e"
	"github.com/guicai123/gin-v2/service/index_service"
	"net/http"
)

func Index(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
	)

	BannerServer := index_service.IndexInfo{
	}
	banners, err := BannerServer.GetBannerOne()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DOLOGIN_FAIL, nil)
		return
	}
	memberInfo := make(map[string]interface{})
	memberInfo["banner1"] = make(map[string]interface{})//空字典
	memberInfo["banner2"] =make([]string, 0)//空数组
	memberInfo["banner3"] = banners
	appG.Response(http.StatusOK, e.SUCCESS, memberInfo)
}
