package main

import (
	"MakeProgress/gin-blog/gin-blog/pkg/setting"
	"MakeProgress/gin-blog/gin-blog/routers"
	"fmt"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,  //允许读取的最大时间
		WriteTimeout:   setting.WriteTimeout, //允许写入的最大时间
		MaxHeaderBytes: 1 << 20,              //请求头的最大字节数
	}
	s.ListenAndServe()
}
