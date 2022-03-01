package storage

import (
	"api/pkg/config"
	"crypto/tls"
	_ "github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis"
)

func RedisStore() *redis.Storage {
	cfg := config.GetConf()
	store := redis.New(redis.Config{
		Host:      cfg.Redis.RedisHost,
		Port:      cfg.Redis.RedisPort,
		Username:  cfg.Redis.RedisUser,
		Password:  cfg.Redis.RedisPass,
		URL:       cfg.Redis.RedisUrl,
		Database:  cfg.Redis.RedisDatabase,
		Reset:     cfg.Redis.RedisReset,
		TLSConfig: &tls.Config{},
	})
	return store
}
