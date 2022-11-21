package service

import (
	"Alpha-S/config"
	"Alpha-S/module/amap"
	"Alpha-S/utils"
	"encoding/json"
	"fmt"
)

// GetWeatherInfo 实时天气
func GetWeatherInfo(envConfig *config.Config) (s string) {
	// 实时天气
	url := fmt.Sprintf("%s/v3/weather/weatherInfo?city=%s&key=%s", envConfig.Amap.Host, "320200", envConfig.Amap.Key)
	bytes, err := utils.SendHttpRequest("GET", url, nil, nil)
	if err != nil {
		fmt.Printf("getWeather error:%s \n", err.Error())
		return
	}
	var weather1 amap.WeatherResp
	err = json.Unmarshal(bytes, &weather1)
	if err != nil {
		fmt.Printf("weatherResp Unmarshal error:%s \n", err.Error())
		return
	}

	s = fmt.Sprintf("今天气温 %s 摄氏度,查看下效果", weather1.Lives[0].Temperature)
	println(s)
	return
}

// GetWeatherForcast 天气预报
func GetWeatherForcast(envConfig *config.Config) (s string) {
	url := fmt.Sprintf("%s/v3/weather/weatherInfo?city=%s&key=%s/all", envConfig.Amap.Host, "320200", envConfig.Amap.Key)
	bytes, err := utils.SendHttpRequest("GET", url, nil, nil)
	if err != nil {
		fmt.Printf("getWeather error:%s \n", err.Error())
		return
	}
	var weather1 amap.WeatherResp
	err = json.Unmarshal(bytes, &weather1)
	if err != nil {
		fmt.Printf("weatherResp Unmarshal error:%s \n", err.Error())
		return
	}

	return
}
