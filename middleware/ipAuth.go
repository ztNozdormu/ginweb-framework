package middleware

import (
	"errors"
	"fmt"
	"github.com/ztNozdormu/ginweb-framework/public"
	"github.com/ztNozdormu/gweb-common/lib"
	"github.com/gin-gonic/gin"
)
/**
 * IP 白名单处理
 */
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range lib.GetStringSliceConf("base.http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched{
			public.ResponseError(c, public.InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
