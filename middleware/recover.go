package middleware

import (
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

func MyRecovery(c *gin.Context) {
	defer func() {
		r := recover()
		if r != nil {
			println("全局异常处理:", r)
			debug.PrintStack()
			//initJob.Error(c, 500, "", r)
			c.Abort()
		}
	}()
	c.Next()
}
