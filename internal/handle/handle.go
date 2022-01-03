package handle

import (
	"e-commerce-api/internal/controller/faq"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {

	// Faq Routes
	app.Post("/faq", faq.Store)
	app.Get("/faq", faq.Index)
	app.Get("/faq/:id", faq.Show)
	app.Post("/faq/delete/:id", faq.Destroy)
	// Faq Routes End
}
