package utils

import (
	"github.com/spf13/viper"
)

var configPath string
var configs map[string]*viper.Viper

func init() {
	configs = make(map[string]*viper.Viper)
	configPath = "config/"
}

func GetConfig(filename string) (*viper.Viper, error) {
	if val, ok := configs[filename]; ok {
		return val, nil
	}

	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(filename)
	v.SetConfigType("json")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	configs[filename] = v
	return v, nil
}

func GetString(filename string, key string) string {
	config, err := GetConfig(filename)
	if err != nil {
		return ""
	}
	if config == nil {
		return ""
	}

	return config.GetString(key)
}
