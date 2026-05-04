package util

import (
	setting "example.com/m/pkg"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	res := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		return (page - 1) * setting.PageSize
	}
	return res
}
