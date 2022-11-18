package router

import (
	"Alpha-S/middleware"
	"Alpha-S/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	// gin的中间件的最终执行顺序类似洋葱模型,c.next前会执行下一层.
	// func(c *gin.Context)及func foo() gin.HandlerFunc都是中间件

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	// 路由组
	// apiV1用于接收 bot
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.RequestFilter(), middleware.ResponseHeader())
	apiV1.POST("/event", service.GetEvent) // 事件上报

	return r
}
