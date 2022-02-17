package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Minio    Minio
	Vault    Vault
	Database Database
}

type Server struct {
	AppName   string
	BodyLimit int
	Port      int
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

func LoadEnvironment() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetConfigFile("config.yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}
func ParseEnvironment(v *viper.Viper) (Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func GetConf() Config {
	var vip, _ = LoadEnvironment()
	var c, _ = ParseEnvironment(vip)
	return c
}
