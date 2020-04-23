package routers

import (
	"MakeProgress/gin-blog/gin-blog/pkg/setting"
	v1 "MakeProgress/gin-blog/gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//Recovery恢复中间件引擎
	//添加一个日志记录器
	r.Use(gin.Logger(), gin.Recovery())
	//设置gin模式
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTag)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新文章标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除文章标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
