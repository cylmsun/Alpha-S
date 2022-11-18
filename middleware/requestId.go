package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func RequestFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			//init.Error(c, http.StatusUnauthorized, "没有授权", "")
			c.Abort()
			return
		}

		uuidStr := uuid.NewV4().String()
		uuidStr = strings.ReplaceAll(uuidStr, "-", "")
		c.Header("X-Request-Id", uuidStr)
		// 设置requestId到context中
		c.Set("request_id", uuidStr)
		c.Next()
	}
}
