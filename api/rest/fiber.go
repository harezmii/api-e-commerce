package rest

import (
	_ "api/docs"
	"api/ent"
	"api/internal/entity/response"
	"api/internal/entity/seed"
	"api/internal/handle"
	"api/internal/secret/middleware"
	_ "api/internal/secret/vault"
	"api/internal/storage"
	"api/pkg/config"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"time"
)

func RunRest(port string) {
	cfg := config.GetConf()

	app := fiber.New(fiber.Config{
		AppName:       cfg.Server.AppName,
		ServerHeader:  cfg.Server.ServerHeader,
		CaseSensitive: cfg.Server.CaseSensitive,
		BodyLimit:     cfg.Server.BodyLimit,
	})
	// Storage
	store := storage.RedisStore()
	// Storage End

	// Cache
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		CacheControl: true,
		Expiration:   time.Second * 8,
		Storage:      store,
		CacheHeader:  "Cache-Time",
	}))
	// Cache End

	// Fiber Internal Middleware
	middleware.SetupMiddleware(app)
	// Fiber Internal Middleware End

	//app.Use(firabaseAuth.FirebaseMiddleWare)

	// Api ping
	app.Get("/", func(ctx *fiber.Ctx) error {
		if ctx.Query("seeder") == "true" {
			seeder := seed.Seeder{
				Client:   ent.ConnectionEnt(),
				Context:  context.Background(),
				SeedInt:  20,
				Entities: []string{"Faq", "Message"}}
			seeder.Seed()
		}
		return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{Message: "Api is Running", StatusCode: fiber.StatusOK, Data: ctx.GetReqHeaders()})
	})

	// Api ping End

	app.Get("swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api")
	version1 := api.Group("/v1")

	handle.SetupRoutes(version1)

	// Match Any Request
	app.Use(func(ctx *fiber.Ctx) error {
		c := ConfigAccept{
			AcceptEncodings: []string{"gzip", "br"},
			AcceptLanguages: []string{"tr", "en"},
			AcceptCharset:   []string{"utf-8", "utf-16"},
			Accepts:         []string{"application/json", "application/xml", "multipart/form-data", "image/png", "image/jpeg", "image/jpg"},
		}
		c.NewAccept(ctx)

		// logs.Logger(ctx, "Any Request!The page you are looking for could not be found.", logs.INFO)
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
