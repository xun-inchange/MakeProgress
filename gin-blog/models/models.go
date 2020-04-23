package models

//用于models的初始化
import (
	"MakeProgress/gin-blog/gin-blog/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database':", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
	))
	if err != nil {
		log.Println(err)
	}
	//设置默认表明前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	//这个配置是单数结构体操作单数表 复数结构体操作复数表。 不加这个配置单数结构体名也能操作复数结构体名的表
	db.SingularTable(true)
	//设置空闲连接池中的最大连接数
	db.DB().SetMaxIdleConns(10)
	//设置数据库连接最大打开数
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
