package main

import (
	"context"
	"github.com/gin-gonic/gin"
    "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
  其他扩展
 */

func main(){
	g:=gin.Default()
	g.GET("/goodClose", func(gin *gin.Context) {
		time.Sleep(10*time.Second)
		gin.String(200,"优雅关停...")
	})
    // 优雅关停服务--在关停前时间内处理完已经接收的请求
	server:=&http.Server{
		Addr:":8088",
		Handler:g,
	}
	// 优雅关停操作通过go程来执行
    go func (){
    	// 监听服务运行状态
        if err:=server.ListenAndServe();err!=nil&&err!=http.ErrServerClosed{
        	log.Fatal("listen:%s\n:",err)
		}
	}()
	// 定义一个信号通道 --系统信号
	quit:=make(chan os.Signal)
	// kill -9无法捕获
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	// 阻塞
	<-quit
	log.Println("shutdown server......")
	// 创建一个超时上下文 收到关停指令后10s再关停
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*10)
    defer cancel()
	if err:=server.Shutdown(ctx);err!=nil{
		log.Fatal("server shutdown:",err)
	}
	log.Println("server exiting...")

}
