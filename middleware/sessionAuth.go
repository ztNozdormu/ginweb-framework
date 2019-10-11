package middleware

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ztNozdormu/ginweb-framework/public"
)
/**
 * 用户是否登录处理
 */
func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if name,ok:=session.Get("user").(string);!ok||name==""{
			public.ResponseError(c, public.InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
