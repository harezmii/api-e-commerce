package rest

import (
	_ "e-commerce-api/docs"
	"e-commerce-api/internal/handle"
	db "e-commerce-api/internal/infraStructure/database"
	_ "e-commerce-api/internal/secret/vault"
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
)

func RestRun(port string) {
	app := fiber.New(fiber.Config{
		AppName:   "E Commerce REST Api",
		BodyLimit: 4096,
	})

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
			return ctx.Status(429).JSON(fiber.Map{
				"statusCode": 429,
				"message":    "Too Many Request",
			})
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
		return ctx.JSON(fiber.Map{
			"statusCode": 200,
			"message":    "Api is running",
		})
	})

	// Api ping End

	app.Get("swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api")
	version1 := api.Group("/v1")

	handle.SetupRoutes(version1)

	serverError := app.Listen("0.0.0.0:" + port)
	if serverError != nil {
		_ = fmt.Sprintf("Server Error")
	}
}
