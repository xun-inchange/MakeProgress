package api

import (
	"MakeProgress/gin-blog/gin-blog/models"
	"MakeProgress/gin-blog/gin-blog/pkg/e"
	"MakeProgress/gin-blog/gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	UserName string `valid:"Required;MaxSize(50)"`
	PassWord string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	userName := c.Query("username")
	passWord := c.Query("password")
	valid := validation.Validation{}

	a := auth{UserName: userName, PassWord: passWord}
	//验证这个结构体
	//这里验证就比较方便，可以直接判断用户输入的符合要求不
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	//struct格式验证通过
	if ok {
		//检查是否存在该用户
		isExist := models.CheckAuth(userName, passWord)
		//如果存在
		if isExist {
			//如果存在就生成一个token
			token, err := util.GenerateToken(userName, passWord)
			if err != nil {
				//生成token错误
				code = e.ERROR_AUTH_TOKEN
			} else {
				//返回的数据
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			//不存在该用户
			code = e.ERROR_AUTH
		}
	} else {
		//struct验证失败
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
