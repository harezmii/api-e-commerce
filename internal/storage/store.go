package storage

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
	"time"
)

func NewMysqlStore(host, database, table, user, pass string, port int, reset bool, time time.Duration) *mysql.Storage {
	store := mysql.New(mysql.Config{
		Host:       host,
		Port:       port,
		Database:   database,
		Username:   user,
		Password:   pass,
		Table:      table,
		Reset:      reset,
		GCInterval: time,
	})
	return store
}

var Storage = NewMysqlStore("127.0.0.1", "storage", "storeSession", "root", "", 3306, true, time.Minute)
var store = session.New(session.Config{Storage: Storage, Expiration: time.Minute})

func GetSession(ctx *fiber.Ctx) interface{} {
	getter, getError := store.Get(ctx)
	if getError != nil {
		return "Get error"
	}
	fmt.Println(getter.Keys())
	return getter.Get("auth")
}
func SetSesion(ctx *fiber.Ctx) bool {
	getter, err := store.Get(ctx)
	if err != nil {
		return false
	}
	getter.Set("auth", true)
	return true
}
