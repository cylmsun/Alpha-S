package middleware

import (
	"github.com/gin-gonic/gin"
)

func MyLogger(c *gin.Context) {
	defer func() {
		//now := time.Now()
		//path := c.Request.URL.Path
		//query := c.Request.URL.RawQuery
		//
		//since := time.Since(now)
		//initJob.Logger.Info(path,
		//	zap.String("method", c.Request.Method),
		//	zap.String("path", path),
		//	zap.String("query", query),
		//	zap.String("IP", c.ClientIP()),
		//	zap.String("user-agent", c.Request.UserAgent()),
		//	zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		//	zap.Duration("duration", since),
		//)
	}()
	c.Next()
}
