package pagination

import (
	"github.com/yunshu2009/blog-service/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()

	if page <=0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	//todo

	return 0
}

func GetPageOffset(page, pageSize int) int {
	// todo

	return 0
}