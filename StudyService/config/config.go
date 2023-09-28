package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./config/config.env")
	err := viper.ReadInConfig()
	if err != nil {
		panic("can not read config file")
	}

	Singleton = &SConfig{}
}

var (
	Singleton Config
)

type Config interface {
	GetConfig(key ConfigKey) interface{}
}

type SConfig struct {
}

func (s *SConfig) GetConfig(key ConfigKey) interface{} {
	return viper.Get(key.key)
}
