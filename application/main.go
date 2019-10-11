package main

import (
	"github.com/ztNozdormu/ginweb-framework/public"
	"github.com/ztNozdormu/ginweb-framework/router"
	"github.com/ztNozdormu/gweb-common/lib"
	"os"
	"os/signal"
	"syscall"
)

func main(){
   // 1.初始化系统配置信息【基本配置:base;数据库配置:mysql;缓存中间件配置:redis...】
	lib.InitModule("./conf/dev/",[]string{"base","mysql","redis",})
	defer lib.Destroy()
	public.InitMysql()
	// 初始化验证器
	public.InitValidate()

	// 启动服务
	router.HttpServerRun()

   // 优雅关停服务
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()
}

