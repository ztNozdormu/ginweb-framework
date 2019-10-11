package public

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ztNozdormu/gweb-common/lib"
)

type ResponseCode int
/**
 * 服务处理响应结果处理
 */
//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode

	InvalidRequestErrorCode ResponseCode = 401
	CustomizeCode           ResponseCode = 1000

	GROUPALL_SAVE_FLOWERROR ResponseCode = 2001
)

type Response struct {
	Code ResponseCode `json:"code"`           // 响应结果的编号
	Msg  string       `json:"msg"`            // 响应结果的描述
	Data      interface{}  `json:"data"`      // 响应结果数据
	TraceId   interface{}  `json:"traceId"`   // 日志追踪ID
	Success bool  `json:"success"`                  // 业务处理结果是否成功
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp := &Response{Code: code, Msg: err.Error(), Data: "", TraceId: traceId,Success:false}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{},msg string) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp := &Response{Code: SuccessCode, Msg:msg, Data: data, TraceId: traceId,Success:true}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
