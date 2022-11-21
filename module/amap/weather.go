package amap

type WeatherResp struct {
	Status   int           `json:"status,string"`
	Count    int           `json:"count,string"`
	Info     string        `json:"info"`
	Lives    []LiveWeather `json:"lives"`
	Forecast Forecast      `json:"forecast"`
}

type WeatherLocation struct {
	Province string `json:"province"`
	City     string `json:"city"`
	Adcode   string `json:"adcode"`
	Humidity string `json:"humidity"`
}

type LiveWeather struct {
	WeatherLocation
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower     string `json:"windpower"`
}

type Forecast struct {
	WeatherLocation
	Casts []Cast `json:"casts"`
}

type Cast struct {
	Date         string `json:"date"`
	Week         int    `json:"week"`
	Dayweather   string `json:"dayweather"`
	Nightweather string `json:"nightweather"`
	Daytemp      string `json:"daytemp"`
	Nighttemp    string `json:"nighttemp"`
	Daywind      string `json:"daywind"`
	Nightwind    string `json:"nightwind"`
	Daypower     string `json:"daypower"`
	Nightpower   string `json:"nightpower"`
}
