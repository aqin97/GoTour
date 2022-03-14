package app

import (
	"github.com/aqin97/GoTour/blog-service/global"
	"github.com/aqin97/GoTour/blog-service/pkg/convert"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pagesize := convert.StrTo(c.Query("page_size")).MustInt()
	if pagesize <= 0 {
		pagesize = global.AppSetting.DefaultPageSize
	}
	if pagesize > global.AppSetting.MaxPageSize {
		pagesize = global.AppSetting.MaxPageSize
	}
	return pagesize
}

func GetPageOffset(page, pagesize int) int {
	var res int
	if page > 0 {
		res = (page - 1) * pagesize
	}
	return res
}
