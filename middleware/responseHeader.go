package middleware

import "github.com/gin-gonic/gin"

func ResponseHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
	}
}
