package main

import "github.com/gin-gonic/gin"

/**
   模板渲染
 */
func main(){
    g:=gin.Default()
    g.LoadHTMLGlob("demo/extend/*.html")
    g.GET("templatePaint", func(gc *gin.Context) {
    	paintParam:=gc.DefaultQuery("paintParam","gin模板渲染的内容")
		gc.HTML(200,"template.html",gin.H{
			"title":"gin模板渲染",
			"content":paintParam,
		})
	})
	g.Run()
}
