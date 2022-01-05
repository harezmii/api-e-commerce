package handle

import (
	"e-commerce-api/internal/controller/faq"
	"e-commerce-api/internal/controller/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {

	// Faq Routes
	app.Post("/faq", faq.Store)
	app.Get("/faq", faq.Index)
	app.Get("/faq/:id", faq.Show)
	app.Delete("/faq/:id", faq.Destroy)
	app.Put("/faq/:id", faq.Update)
	// Faq Routes End

	// User Routes
	app.Post("/user", user.Store)
	app.Get("/user", user.Index)
	app.Get("/user/:id", user.Show)
	app.Delete("/user/:id", user.Destroy)
	app.Put("/user/:id", user.Update)
	// User Routes End

}
