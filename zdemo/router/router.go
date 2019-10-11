package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

func main(){
	grouter:=gin.Default()

	//grouter.GET("/ginGET", func(gin *gin.Context) {
	//	gin.String(200,"ginGET")
	//})
	grouter.POST("/ginPOST", func(gin *gin.Context) {
		gin.String(200,"ginPOST")
	})
	grouter.PUT("/ginPUT", func(gin *gin.Context) {
		gin.String(200,"ginPUT")
	})
	grouter.PATCH("/ginPATCH", func(gin *gin.Context) {
		gin.String(200,"ginPATCH")
	})
	grouter.DELETE("/ginDELETE", func(gin *gin.Context) {
		gin.String(200,"ginDELETE")
	})
	//// 支持八种请求类型
	//grouter.Any("/ginAny", func(gin *gin.Context) {
	//	gin.String(200,"ginAny")
	//})
	//// 绑定静态文件
	//grouter.Static("/assets","/assets")
	//grouter.StaticFS("/static",http.Dir("static"))
	//grouter.StaticFile("favicon.ico","./favicon.ico")
	// 参数作为URL 注意该方法会与上面有同类型的冲突
	//grouter.GET("/:name/:id", func(g *gin.Context) {
	//	g.JSON(200,gin.H{
	//		"name":g.Param("name"),
	//		"id":g.Param("id"),
	//	})
	//})
	// 泛绑定  前缀为 user的请求都会接收
	//grouter.GET("/user/*action", func(gin *gin.Context) {
	//	gin.String(200,"泛绑定")
	//})
	// get参数获取
	//grouter.GET("/getParam", func(g *gin.Context) {
	//	firstName :=g.Query("firstName")
	//	lastName:=g.DefaultQuery("lastName","lastDefaultName")
	//	g.JSON(200,gin.H{
	//		"firstNam":firstName,
	//		"lastName":lastName,
	//	})
	//})
	// body参数获取 --前台一般传递Json数据格式
	//grouter.GET("/getBody", func(g *gin.Context) {
	//	bodyBytes,err:=ioutil.ReadAll(g.Request.Body)
	//	if err!=nil{
	//		g.String(http.StatusBadRequest,err.Error())
	//		g.Abort()
	//	}
	//	// firstName=xxx&&lastName=xxx需要处理数据后重新设置到g.Request.Body再获取才能获取得到
	//	g.Request.Body=ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	//	firstName :=g.PostForm("firstName")
	//	lastName:=g.DefaultPostForm("lastName","default_last_name")
	//	g.String(http.StatusOK,"s%,s%,s%",firstName,lastName,string(bodyBytes))
	//})
	//  获取bind参数 参数自动绑定到对应的结构体
	//  参数验证 1.结构体验证 【通过tag进行验证】 2.自定义验证 3.升级验证 支持多语言错误信息
	grouter.GET("bindGET",bindParam)
	grouter.POST("bindPOST",bindParam)

    // 自定义验证规则
	grouter.GET("currentCheckParam",currentCheckParam)

	grouter.Run()
}
func currentCheckParam(c *gin.Context){
	// 通过断言返回一个Validator 验证器
	if v,ok:=binding.Validator.Engine().(*validator.Validate);ok{
		// 将自定义的规则customCheckStart注册到验证器中，key value和取的名字一样
		v.RegisterValidation("customCheckStart",customCheckStart)
	}
	var custmonCheck CustmonCheck
	// 这里是根据不同的content-type做不同的绑定处理
	if err:=c.ShouldBind(&custmonCheck);err==nil{
		c.String(200,"%v",custmonCheck)
	}else{
		c.String(500,"pserson bind err:%v",err)
		c.Abort()
		return
	}
}
// 自定义校验规则 结构体DEMO
type CustmonCheck struct {
	// customCheckStart：自定义的校验规则函数名称
	CheckStart time.Time `form:"checkStart" binding:"required,customCheckStart" time_format:"2006-01-02"`
	// gtfield=CheckStart:字段CheckEnd的值必须大于CheckStart字段的值
	CheckEnd time.Time `form:"checkEnd" binding:"required,gtfield=CheckStart" time_format:"2006-01-02"`
}

// 自定义的验证规则
func customCheckStart(v *validator.Validate,topStruct reflect.Value,
	currentStructOrField reflect.Value,field reflect.Value,fieldType reflect.Type,
	fieldKind reflect.Kind,param string) bool{
		if date,ok:=field.Interface().(time.Time);ok{
			currentTime:=time.Now()
			// 规则:被验证参数大于当前时间
			if date.Unix()>currentTime.Unix(){
             return true
			}
		}
		return false
}

type Person struct {
	// binding规则:多验证条件与关系用,逗号隔开，或关系用| 竖线隔开；更详细的规则查看文档
	Age int     `form:"age" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}
func bindParam(c *gin.Context){
	var pserson Person
	// 这里是根据不同的content-type做不同的绑定处理
	if err:=c.ShouldBind(&pserson);err==nil{
        c.String(200,"%v",pserson)
	}else{
		c.String(500,"pserson bind err:%v",err)
		c.Abort()
		return
	}
}
