package util

//工具包

import (
	"MakeProgress/gin-blog/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//编写分页页码的获取方法
func GetPage(c *gin.Context) int {
	result := 0
	//StrTo()是一个封装好的包 里面有各种将字符串转移成其它类型的方法
	//page 为页数
	page, _ := com.StrTo(c.Query("page")).Int()
	//如果页数大于0了
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
