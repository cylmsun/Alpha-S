package init

import (
	"Alpha-S/config"
	"fmt"
	"github.com/robfig/cron/v3"
)

func Cron(envConfig config.Config) {
	c := cron.New()
	println("准备cron")

	// cron/v3 按照linux标准只有5位，6位会err
	_, err := c.AddFunc(envConfig.Server.Cron, job)
	if err != nil {
		fmt.Printf("cron set error:%s \n", err.Error())
		return
	}

	c.Start()
	return
}

func job() {
	fmt.Println("job run...")
}
