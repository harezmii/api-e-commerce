package storage

//
//
//import (
//	"fmt"
//	"github.com/gofiber/fiber/v2"
//	"github.com/gofiber/fiber/v2/middleware/session"
//	"github.com/gofiber/storage/mysql"
//	"time"
//)
//
//func NewMysqlStore(host, database, table, user, pass string, port int, reset bool, time time.Duration) *mysql.Storage {
//	store := mysql.New(mysql.Config{
//		Host:       host,
//		Port:       port,
//		Database:   database,
//		Username:   user,
//		Password:   pass,
//		Table:      table,
//		Reset:      reset,
//		GCInterval: time,
//	})
//	return store
//}
//
//var Storage = NewMysqlStore("127.0.0.1", "storage", "storeSession", "root", "", 3306, true, time.Minute)
//var store = session.New(session.Config{Storage: Storage, Expiration: time.Minute})
//
//func GetSession(ctx *fiber.Ctx) interface{} {
//	getter, getError := store.Get(ctx)
//	if getError != nil {
//		return "Get error"
//	}
//	fmt.Println(getter.Keys())
//	return getter.Get("auth")
//}
//func SetSesion(ctx *fiber.Ctx) bool {
//	getter, err := store.Get(ctx)
//	if err != nil {
//		return false
//	}
//	getter.Set("auth", true)
//	return true
//}
import (
	"api/pkg/config"
	_ "github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis"
)

func RedisStore() *redis.Storage {
	store := redis.New(redis.Config{
		Host:      config.GetEnvironment("REDIS_HOST", config.STRING).(string),
		Port:      config.GetEnvironment("REDIS_PORT", config.INTEGER).(int),
		Username:  config.GetEnvironment("REDIS_USER", config.STRING).(string),
		Password:  config.GetEnvironment("REDIS_PASS", config.STRING).(string),
		URL:       config.GetEnvironment("REDIS_URL", config.STRING).(string),
		Database:  config.GetEnvironment("REDIS_DATABASE", config.INTEGER).(int),
		Reset:     config.GetEnvironment("REDIS_RESET", config.BOOL).(bool),
		TLSConfig: nil,
	})
	return store
}
