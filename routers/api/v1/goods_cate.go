package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guicai123/gin-v2/models"
	"github.com/guicai123/gin-v2/pkg/app"
	"github.com/guicai123/gin-v2/pkg/e"
	"github.com/guicai123/gin-v2/pkg/setting"
	"github.com/guicai123/gin-v2/pkg/util"
	"math"
	"net/http"
)

func GoodsCate(c *gin.Context) {
	appG := app.Gin{C: c}
	goods_cate := map[string]interface{}{
		"is_deleted": 0,
	}
	list, _ := models.GetGoodsCate(util.GetPage(c), setting.AppSetting.PageSize, goods_cate)
	count, _ := models.GetGoodsCateTotal(goods_cate)
	data := make(map[string]interface{})
	data["lists"] = list
	data["pagecount"] = math.Ceil(float64(count) / float64(setting.AppSetting.PageSize))
	appG.Response(http.StatusOK, e.SUCCESS, data)

}

func GoodsTCate(c *gin.Context) {
	appG := app.Gin{C: c}

	list, _ := models.GetGoodsInfo()
	data := make(map[string]interface{})
	data["lists2"] = list

	appG.Response(http.StatusOK, e.SUCCESS, data)

}
