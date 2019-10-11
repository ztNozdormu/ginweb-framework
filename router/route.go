package router

import (

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ztNozdormu/ginweb-framework/controller"
	"github.com/ztNozdormu/ginweb-framework/middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	//写入gin日志
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.DefaultErrorWriter = io.MultiWriter(f)
	router := gin.Default()
	router.Use(middlewares...)

	//demo
	v1 := router.Group("/demo")
	v1.Use(middleware.RecoveryMiddleware(), middleware.RequestLog(), middleware.IPAuthMiddleware(), middleware.TranslationMiddleware())
	{
		controller.DemoRegister(v1)
	}

	//api
	store := sessions.NewCookieStore([]byte("secret"))
	// 创建通用路由组 【包含api的请求地址都在该路由组下】
	apiNormalGroup := router.Group("/api")
	apiController:=&controller.Api{}
	// 路由组使用的中间件
	apiNormalGroup.Use(
		// 设置session
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	// 指定子路由
	apiNormalGroup.POST("/login",apiController.Login)
	apiNormalGroup.GET("/loginout",apiController.LoginOut)

	// API权限组 --除了登录和登出不需要验证session.其他接口都需要做session【基本要求】，以及根据用户做相关权限校验
	apiAuthGroup := router.Group("/api")
	apiAuthGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		// session校验中间件
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	// 指定子路由
	apiAuthGroup.GET("/user/listpage", apiController.ListPage)
	apiAuthGroup.GET("/user/add", apiController.AddUser)
	apiAuthGroup.GET("/user/edit", apiController.EditUser)
	apiAuthGroup.GET("/user/remove", apiController.RemoveUser)
	apiAuthGroup.GET("/user/batchremove", apiController.RemoveUser)
	return router
}
