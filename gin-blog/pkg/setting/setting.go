package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

//设置配置参数
var (
	Cfg          *ini.File //内存中一个或多个ini文件的组合
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration //纳秒(时长,耗时)
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	//加载解析ini数据源
	Cfg, err = ini.Load("gin-blog/conf/app.ini")
	if err != nil {
		log.Fatal("Fail to parse conf/app.ini:", err)
	}
	LoadBase()
	LoadApp()
	LoadServer()
}

func LoadBase() {
	//读取操作，默认分区可以用空字符表示
	//MustString：传递默认值,在读取key为RUN_MODE的时候，如果为空就传递默认值
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	//获取指定分区
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get section 'server':", err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal("Fail to get section 'app':", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
