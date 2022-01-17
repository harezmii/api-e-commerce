package config

import (
	"github.com/spf13/viper"
)

const (
	INTEGER int = 0
	STRING  int = 1
	BOOL    int = 2
	FLOAT   int = 3
	MAP     int = 4
)

func GetEnvironment(key string, process int) interface{} {
	if loadEnvironment() {
		switch process {
		case INTEGER:
			{
				return viper.GetInt(key)
				break
			}
		case STRING:
			{
				return viper.GetString(key)
			}
		case BOOL:
			{
				return viper.GetBool(key)
			}
		case FLOAT:
			{
				return viper.GetFloat64(key)
			}
		case MAP:
			{
				return viper.GetStringMap(key)
			}

		}
	}
	return "Get environment error"
}

func loadEnvironment() bool {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return false
	}
	return true
}
