package rest

import (
	_ "api/docs"
	"api/internal/entity/response"
	"api/internal/handle"
	db "api/internal/infraStructure/database"
	_ "api/internal/secret/vault"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/storage/mysql"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"time"
)

func RestRun(port string) {
	app := fiber.New(fiber.Config{
		AppName:   "E Commerce REST Api",
		BodyLimit: 4096,
	})

	// Storage
	store := mysql.New(mysql.Config{
		Host:       "127.0.0.1",
		Port:       3306,
		Database:   "storage",
		Username:   "root",
		Password:   "",
		Table:      "store",
		Reset:      false,
		GCInterval: 3 * time.Second,
	})
	// Storage End

	// Cache
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		CacheControl: true,
		Storage:      store,
		CacheHeader:  "Cache-Time",
	}))

	// Cache End

	// Logger

	app.Use(logger.New(logger.Config{}))
	// Logger End

	// Database
	defer func() {
		db.PrismaDisConnection()
	}()
	// Database End

	// Fiber Internal Middleware
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(compress.New())
	app.Use(limiter.New(limiter.Config{
		Max: 10,
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(429).JSON(response.ErrorResponse{StatusCode: 429, Message: "Too many Request"})
		},
	}))
	app.Use(requestid.New())
	app.Use(recover2.New())
	app.Use(favicon.New())

	// Fiber Internal Middleware End

	// Helmet
	app.Use(helmet.New())

	// Helmet End

	//app.Use(firabaseAuth.FirebaseMiddleWare)

	// Api ping
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(response.SuccessResponse{Message: "Api is Running", StatusCode: 200})
	})

	// Api ping End

	app.Get("swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api")
	version1 := api.Group("/v1")

	handle.SetupRoutes(version1)

	// Match Any Request
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(404).JSON(response.ErrorResponse{
			StatusCode: 404,
			Message:    "The page you are looking for could not be found.",
		})
	})
	serverError := app.Listen("0.0.0.0:" + port)
	if serverError != nil {
		_ = fmt.Sprintf("Server Error")
	}
}
