package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreateBy   string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
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
