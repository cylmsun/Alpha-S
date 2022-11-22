package service

import (
	"Alpha-S/module/amap"
	"Alpha-S/utils"
	"encoding/json"
	"fmt"
)

// GetWeatherInfo 实时天气
// todo 匹配根据城市查询
func GetWeatherInfo(cityCode string) (s string) {
	// 实时天气
	url := fmt.Sprintf("%s/v3/weather/weatherInfo?city=%s&key=%s", utils.CONFIG.Amap.Host, cityCode, utils.CONFIG.Amap.Key)
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

	s = fmt.Sprintf("今天气温 %s 摄氏度，风向 %s， 风力 %s", weather1.Lives[0].Temperature, weather1.Lives[0].Winddirection, weather1.Lives[0].Windpower)
	fmt.Printf("实时天气结果:%s", s)
	return
}

// GetWeatherForcast 天气预报
func GetWeatherForcast() (cityCode string) {
	url := fmt.Sprintf("%s/v3/weather/weatherInfo?city=%s&key=%s/all", utils.CONFIG.Amap.Host, cityCode, utils.CONFIG.Amap.Key)
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
