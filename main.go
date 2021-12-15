package main

import (
	"example/pkg/cron"
	"example/routers"
	"net/http"
)
func main(){
	//引入gin路由
	r:=routers.InitRouter()
	//创建http服务
	s:=&http.Server{
		Addr: ":8080",
		Handler: r,
	}
	cron.Cron()
	s.ListenAndServe()

}
