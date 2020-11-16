package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yunshu2009/blog-service/globalconvert"
	"github.com/yunshu2009/blog-service/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()

	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()

	if pageSize <= 0 {
		return 1
	}

	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	// todo

	return 0
}
