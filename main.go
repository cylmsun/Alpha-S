package main

import (
	"Alpha-S/config"
	"Alpha-S/middleware"
	"Alpha-S/router"
	"Alpha-S/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	envConfig := config.GetConfig()
	utils.CONFIG = &envConfig

	//initJob.Cron(&envConfig)

	r := gin.New()
	r.Use(middleware.MyLogger, middleware.MyRecovery)
	r = router.InitRouter(r)

	err := r.Run()
	if err != nil {
		fmt.Printf("gin run error:%s \n", err.Error())
		_ = envConfig.DBConfig
		return
	}

	// \n 或者 %0a 换行？
}
