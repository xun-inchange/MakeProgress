package v1

import (
	"MakeProgress/gin-blog/gin-blog/models"
	"MakeProgress/gin-blog/gin-blog/pkg/e"
	"MakeProgress/gin-blog/gin-blog/pkg/setting"
	"MakeProgress/gin-blog/gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

//获取多个文章标签
func GetTag(c *gin.Context) {
	//c.Query()用于获取?name=test&state=1这类url参数，c.DefaultQuery()则支持设置一个默认参数
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	//指定code码
	code := e.SUCCESS

	//获取标签列表
	//util.GetPage保证了各接口的page处理是一致的
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	//获取总数
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	//验证数据是否符合规则
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	//如果有错误信息，说明验证没有通过
	//表明验证通过了
	if !valid.HasErrors() {
		//如果不存在这个文章标签就创建
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			//code 返回为已经存在这个标签了
			code = e.ERROR_EXIST_TAG
		}
	} else {
		//如果有错误信息，就遍历出所有的错误.因为Errors是一个切片
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	//添加参数验证
	valid.Min(id, 1, "id").Message("id必须大于0")

	code := e.INVALID_PARAMS
	//如果参数没有错误
	if !valid.HasErrors() {
		code = e.SUCCESS
		//如果存在这个id就删除
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			//返回不存在这个标签
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		//如果参数有错误就把有错误的参数输出到日志上
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改tag
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("id不能为空")
	valid.Required(modifiedBy, "modified_by").Message("modified_by不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
