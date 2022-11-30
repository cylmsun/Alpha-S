package service

import (
	"Alpha-S/module"
	"Alpha-S/utils"
	"fmt"
	"regexp"
	"time"
)

func PrivateLogic(msg string) (replyMsg module.LogicResult) {
	c := make(chan module.LogicResult)
	defer close(c)
	go weatherMatch(msg, c)
	go exchangeRate(msg, c)

	select {
	case replyMsg = <-c:
		return
	case <-time.After(time.Second * 5):
		fmt.Println("没有匹配到")
		return
	}
}

func weatherMatch(msg string, c chan<- module.LogicResult) {
	weatherRegexp := regexp.MustCompile(`^查(.*)天气$`)
	subMatch := weatherRegexp.FindStringSubmatch(msg)
	if len(subMatch) != 2 {
		return
	}
	if subMatch[1] == "" {
		subMatch[1] = "无锡市"
	}
	m := make(map[string]string)
	m["无锡市"] = "320281"
	m["无锡"] = "320200"
	m["江阴"] = "320281"
	m["新吴区"] = "320281"
	if value, ok := m[subMatch[1]]; ok {
		info := GetWeatherInfo(value, m[subMatch[1]])
		image := utils.Text2Image(info)
		lr := module.LogicResult{Result: image, ResultType: 2}
		c <- lr
	}
}

func exchangeRate(msg string, c chan<- module.LogicResult) {
	reg := regexp.MustCompile(`^查(.*)汇率$`)
	subMatch := reg.FindStringSubmatch(msg)

	var s string
	m := make(map[string]string)
	m["美元人民币"] = "USD/CNY"
	m["日元"] = "100JPY/CNY"
	if subMatch == nil {
		return
	}
	switch len(subMatch) {
	case 1:
		s = m["美元人民币"]
	case 2:
		s = m[subMatch[1]]
	default:
		s = m["美元人民币"]
	}
	info := GetExchangeRate(s)
	lr := module.LogicResult{Result: info, ResultType: 1}
	c <- lr
}
