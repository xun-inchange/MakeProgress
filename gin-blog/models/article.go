package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

//文章
//这里利用TagID与Tag关联起来，相当于将Tag与Article关联起来
type Article struct {
	Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//回调函数，
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

//通过ID判断这个文章是否存在
func ExistArticleByID(id int) bool {
	var article Article //接受值
	db.Select("id").Where("id=?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	//db.model()指定模型
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

//获取所有文章
//pageNum:第几页
//pageSize:每页多少条
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

//获取一篇文章
func GetArticle(id int) (article Article) {
	db.Where("id=?", id).First(&article)
	//Related：关联查询
	db.Model(&article).Related(&article.Tag)

	return
}

//更新文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Updates(data)
	return true
}

//增加文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

//删除文章
func DeleteArticle(id int) bool {
	//delete 指定删除哪个表的
	db.Where("id=?", id).Delete(Article{})
	return true
}
