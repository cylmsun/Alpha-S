package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		AppName     string `yaml:"app-name"`
		AccessToken string `yaml:"access-token"`
		Cron        string `yaml:"cron"`
	} `yaml:"server"`
	FrontServer struct {
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		AccessToken string `yaml:"access-token"`
	} `yaml:"front-server"`
	DBConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbName"`
	} `yaml:"db"`
	Amap struct {
		Host string `yaml:"host"`
		Key  string `yaml:"key"`
	} `yaml:"amap"`
	OtherApi struct {
		ExchangeRate string `yaml:"exchangeRate"`
	}
}

// GetConfig 读取配置文件,path是项目根目录开始的相对路径
func GetConfig() Config {
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Printf("read config file error:%s \n", err.Error())
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("config file yaml.Unmarshal error:%s \n", err.Error())
	}
	return config
}
