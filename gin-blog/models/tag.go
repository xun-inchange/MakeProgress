package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`        //标签名
	CreatedBy  string `json:"created_by"`  //创建人
	ModifiedBy string `json:"modified_by"` //修改人
	State      int    `json:"state"`       //状态
}

//可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用
// 如果任何回调返回错误，gorm将停止未来操作并回滚所有更改。
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	//time.Now().Unix()获取当前时间的时间戳,单位为秒
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	//offset() 偏移，也就是跳过几条记录 从第n+1条开始
	//limit设置检索的条数 limit 5:意思就是检索前5条记录,相当于limit 0,5
	//Find() 获取所有记录,然后将值传入到tags里面去
	//where 查询条件
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	//Model() 指定数据库要运行的模型就是对应数据库的哪张表
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

//通过name查询tag.id
//通过标签名判断标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	})
	return true
}

//通过标签id来判断tag是否存在
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//通过标签id来删除标签
func DeleteTag(id int) bool {
	db.Where("id=?", id).Delete(&Tag{})
	return true
}

//编辑Tag
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id=?", id).Updates(data)
	return true
}
