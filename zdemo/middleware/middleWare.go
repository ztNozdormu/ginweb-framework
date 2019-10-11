package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)
/**
  中间件
 */
func main(){
	log,_:=os.Create("demo/middleware/gin.log")//../
	errLog,_:=os.Create("demo/middleware/err.log")//../demo/middleware/
	// 正常\错误日志输入到指定文件
	gin.DefaultWriter=io.MultiWriter(log)
	gin.DefaultErrorWriter=io.MultiWriter(errLog)
	g:=gin.New()
	// gin.Recovery()发生异常服务不中断
	g.Use(IPAuthCheckMiddleWare(),gin.Logger(),gin.Recovery())
	g.GET("/logGET",func(g *gin.Context){
		param:=g.DefaultQuery("param","请求日志记录")
		//panic("发生错误服务不中断")
		g.JSON(200,gin.H{
			"message":"logGET请求成功，日志记录成功!参数是:"+param,
			})

	})
	g.Run()
}
// 自定义中间件  IP白名单 DEMO
func IPAuthCheckMiddleWare() gin.HandlerFunc{
  return func(g *gin.Context){
	  flag := false
  	// 允许访问的IP白名单列表 ---正常开发通过配置文件配置【或者通过其他存储方式存储】
    allowClientIP:=[]string{
    	"127.0.0.1",
	}
  	// 获取当前请求IP地址
  	currentClientIp:=g.ClientIP()

  	for _,ip:=range allowClientIP{
  		if currentClientIp==ip{
			flag = true
			break
	    }
  	}
  	if !flag{
  		g.String(401,"%s,not in whiteIpList:",currentClientIp)
	    g.Abort()
  	}
  }
}
