package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Minio    Minio
	Vault    Vault
	Database Database
	Kafka    Kafka
	Redis    Redis
}

type Server struct {
	AppName       string
	BodyLimit     int
	Port          int
	ServerHeader  string
	CaseSensitive bool
}
type Database struct {
	DatabaseUrl    string
	DriverName     string
	DataSourceName string
}
type Vault struct {
	VaultToken   string
	VaultAddress string
	VaultPath    string
}
type Minio struct {
	EndPoint  string
	AccessKey string
	SecretKey string
}

type Kafka struct {
	Topic string
	Conf  map[string]interface{}
}
type Redis struct {
	RedisHost     string
	RedisPort     int
	RedisUser     string
	RedisPass     string
	RedisUrl      string
	RedisDatabase int
	RedisReset    bool
}

func loadEnvironment() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetConfigFile("config.yaml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}
func parseEnvironment(v *viper.Viper) (Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func GetConf() Config {
	var vip, _ = loadEnvironment()
	var c, _ = parseEnvironment(vip)
	return c
}
