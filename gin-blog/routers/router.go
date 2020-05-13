package routers

import (
	"MakeProgress/gin-blog/gin-blog/middleware/jwt"
	"MakeProgress/gin-blog/gin-blog/pkg/setting"
	"MakeProgress/gin-blog/gin-blog/routers/api"
	v1 "MakeProgress/gin-blog/gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//Recovery恢复中间件引擎
	//添加一个日志记录器
	r.Use(gin.Logger(), gin.Recovery())
	//设置gin具备调试模式
	gin.SetMode(setting.RunMode)
	//增加一个路由,用于用户的登录
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	//加入中间件验签,防止中间件任意访问的问题
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTag)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新文章标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除文章标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/article/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/article/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
