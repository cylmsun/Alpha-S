package service

import (
	"regexp"
)

func PrivateLogic(msg string) (replyMsg string) {
	c := make(chan string, 1)
	defer close(c)
	go WeatherMatch(msg, c)

	select {
	case replyMsg = <-c:
		return
	}
}

func WeatherMatch(msg string, c chan<- string) {
	weatherRegexp := regexp.MustCompile(`^查(.*)天气$`)
	submatch := weatherRegexp.FindStringSubmatch(msg)
	if len(submatch) != 2 {
		return
	}
	var m map[string]string
	m["无锡市"] = "320201"
	m["无锡"] = "320200"
	m["江阴"] = "320281"
	m["新吴区"] = "320281"
	if value, ok := m[submatch[1]]; ok {
		info := GetWeatherInfo(value)
		c <- info
	}
}

func ExchangeRate(msg string, c chan<- string) {
	reg := regexp.MustCompile(`^查(.*)汇率$`)
	submatch := reg.FindStringSubmatch(msg)

	var s string
	switch len(submatch) {
	case 1:
		s = "美元"
	case 2:
		s = submatch[1]
	default:
		s = "美元"
	}
	info := GetExchangeRate(s)
	c <- info
}
