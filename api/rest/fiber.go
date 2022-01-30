package rest

import (
	_ "api/docs"
	"api/internal/entity/response"
	"api/internal/handle"
	prisma "api/internal/infraStructure/prismaClient"
	"api/internal/logs"
	_ "api/internal/secret/vault"
	"api/pkg/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.elastic.co/apm/module/apmfiber"
)

func RestRun(port string) {
	app := fiber.New(fiber.Config{
		AppName:   config.GetEnvironment("APP_NAME", config.STRING).(string),
		BodyLimit: config.GetEnvironment("BODY_LİMİT", config.INTEGER).(int),
	})

	// Storage
	//store := storage.RedisStore()
	// Storage End

	// Cache
	//app.Use(cache.New(cache.Config{
	//	Next: func(c *fiber.Ctx) bool {
	//		return c.Query("refresh") == "true"
	//	},
	//	CacheControl: true,
	//	Storage:      store,
	//	CacheHeader:  "Cache-Time",
	//}))
	// Cache End

	// APM Middleware
	app.Use(apmfiber.Middleware())
	// APM Middleware END

	// Logger
	//app.Use(logger.New())
	// Logger End

	// Database
	defer func() {
		prisma.PrismaDisConnection()
	}()
	// Database End

	// Fiber Internal Middleware
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(compress.New())
	app.Use(limiter.New(limiter.Config{
		Max: 10,
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{StatusCode: 429, Message: "Too many Request"})
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
		logs.Logger(ctx, "Any Request!The page you are looking for could not be found.", logs.INFO)
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
