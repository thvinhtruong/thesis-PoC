package config

import (
	"github.com/spf13/viper"
	"sync"
)

func init() {
	viper.SetConfigFile("./config/config.env")
	viper.ReadInConfig()

	singleton = &SConfig{}
}

var (
	once      sync.Once
	singleton *SConfig
)

type Config interface {
	GetConfig(key ConfigKey) interface{}
}

type SConfig struct {
}

func (s *SConfig) GetConfig(key ConfigKey) interface{} {
	return viper.Get(key.key)
}

func GetInstance() *SConfig {
	return singleton
}
