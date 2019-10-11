package main

import(
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)
func main(){
	g:=gin.Default()
	g.GET("autotlsDown", func(context *gin.Context) {
		context.String(200,"证书自动认证完成")
	})
	autotls.Run(g,"www.baidu.com") // 需要正式外网服务器
}
