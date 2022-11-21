package initJob

import (
	"Alpha-S/config"
	"Alpha-S/service"
	"fmt"
	"github.com/robfig/cron/v3"
)

func Cron(envConfig *config.Config) {
	c := newWithSeconds()

	_, err := c.AddFunc(envConfig.Server.Cron, func() {
		jobTest()
		service.GetWeatherInfo(envConfig)
	})
	if err != nil {
		fmt.Printf("cron set error:%s \n", err.Error())
		return
	}

	c.Start()
	return
}

func jobTest() {
	println("job run...")
}

// robfig/cron 按照linux标准只有5位，6位会err，需要修改下
func newWithSeconds() *cron.Cron {
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(parser), cron.WithChain())
}
