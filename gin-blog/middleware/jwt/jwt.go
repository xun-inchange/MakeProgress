package jwt

import (
	"MakeProgress/gin-blog/gin-blog/pkg/e"
	"MakeProgress/gin-blog/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//Gin的中间件HandleFunc
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		tokenString := c.Query("token")
		if tokenString == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(tokenString)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			//用户没有权限
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			//阻止执行后面的HandleFunc
			c.Abort()
			return
		}
		//继续执行下面的HandleFunc
		c.Next()
	}
}
