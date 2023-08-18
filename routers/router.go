package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/guicai123/gin-v2/pkg/util"
	"github.com/guicai123/gin-v2/routers/api"
	v1 "github.com/guicai123/gin-v2/routers/api/v1"

	"github.com/guicai123/gin-v2/middleware/jwt"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(util.LoggerToFile())
	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/reg", v1.DoReg)     //用户注册
		apiv1.POST("/login", v1.DoLogin) //用户登陆
		apiv1.GET("/index", v1.Index)    //app首页

		apiv1.POST("/tags", v1.AddTag) //添加标题
		apiv1.GET("/tags", v1.GetTags) //获取标签列表

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)

		apiv1.GET("/goodscate", v1.GoodsCate)
		apiv1.GET("/goods-cate", v1.GoodsTCate)

	}

	apiv1.Use(jwt.JWT())
	{
		////获取标签列表
		//apiv1.GET("/tags", v1.GetTags)
		////新建标签
		//apiv1.POST("/tags", v1.AddTag)
		////更新指定标签
		//apiv1.PUT("/tags/:id", v1.EditTag)
		////删除指定标签
		//apiv1.DELETE("/tags/:id", v1.DeleteTag)

		////导出标签
		//r.POST("/tags/export", v1.ExportTag)
		////导入标签
		//r.POST("/tags/import", v1.ImportTag)
		//
		////获取文章列表
		//apiv1.GET("/articles", v1.GetArticles)
		////获取指定文章
		//apiv1.GET("/articles/:id", v1.GetArticle)
		////新建文章
		//apiv1.POST("/articles", v1.AddArticle)
		////更新指定文章
		//apiv1.PUT("/articles/:id", v1.EditArticle)
		////删除指定文章
		//apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		////生成文章海报
		//apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
