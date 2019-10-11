package controller

import (
	"encoding/json"
	"github.com/ztNozdormu/gweb-common/lib"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/ztNozdormu/ginweb-framework/dao"
	"github.com/ztNozdormu/ginweb-framework/dto"
	"github.com/ztNozdormu/ginweb-framework/public"
)

type Demo struct {
}

func DemoRegister(router *gin.RouterGroup) {
	demo := Demo{}
	router.GET("/index", demo.Index)
	router.GET("/bind", demo.Bind)
	router.GET("/dao", demo.Dao)
	router.GET("/redis", demo.Redis)
}

func (demo *Demo) Index(c *gin.Context) {
	public.ResponseSuccess(c, "","")
	return
}

func (demo *Demo) Dao(c *gin.Context) {
	if area,err:=(&dao.Area{}).Find(c.DefaultQuery("id","1"));err!=nil{
		public.ResponseError(c,501,err)
	}else{
		js,_:=json.Marshal(area)
		public.ResponseSuccess(c, string(js),"")
	}
	return
}

func (demo *Demo) Redis(c *gin.Context) {
	redisKey:="redis_key"
	lib.RedisConfDo(public.GetTraceContext(c),"default",
		"SET",redisKey,"redis_value")
	redisValue,err:=redis.String(lib.RedisConfDo(public.GetTraceContext(c),"default",
		"GET",redisKey))
	if err!=nil{
		public.ResponseError(c,501,err)
		return
	}
	public.ResponseSuccess(c, redisValue,"")
	return
}

func (demo *Demo) Bind(c *gin.Context) {
	st:=&dto.InStruct{}
	if err:=st.BindingValidParams(c);err!=nil{
		public.ResponseError(c,500,err)
		return
	}
	public.ResponseSuccess(c, "","")
	return
}
