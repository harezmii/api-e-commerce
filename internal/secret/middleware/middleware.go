package middleware

import (
	"api/internal/entity/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"go.elastic.co/apm/module/apmfiber"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(apmfiber.Middleware())
	app.Use(helmet.New())
	app.Use(recover2.New())
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(compress.New())
	app.Use(limiter.New(limiter.Config{
		Max: 120,
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{StatusCode: 429, Message: "Too many Request"})
		},
	}))
	app.Use(requestid.New())
	app.Use(recover2.New())
	app.Use(favicon.New())
}
