package routers

import (
	"example.com/m/middleware/jwt"
	setting "example.com/m/pkg"
	v1 "example.com/m/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", v1.GetAuth)
	v1Group := r.Group("api/v1")

	v1Group.Use(jwt.JWT())
	{
		//获取标签列表
		v1Group.GET("/tags", v1.GetTags)
		//新建标签
		v1Group.POST("/tags", v1.AddTag)
		//更新指定标签
		v1Group.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		v1Group.DELETE("/tags/:id", v1.DeleteTag)

		// 获取文章列表
		v1Group.GET("/articles", v1.GetArticles)
		// 获取指定文章
		v1Group.GET("/articles/:id", v1.GetArticle)
		// 新建文章
		v1Group.POST("/articles", v1.AddArticle)
		// 更新指定文章
		v1Group.PUT("/articles/:id", v1.EditArticle)
		// 删除指定文章
		v1Group.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
