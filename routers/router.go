package routers

import (
	setting "example.com/m/pkg"
	v1 "example.com/m/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	v1Group := r.Group("api/v1")
	{
		//获取标签列表
		v1Group.GET("/tags", v1.GetTags)
		//新建标签
		v1Group.POST("/tags", v1.AddTag)
		//更新指定标签
		v1Group.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		v1Group.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
