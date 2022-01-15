package handle

import (
	"api/internal/controller/faq"
	"api/internal/controller/profile"
	"api/internal/controller/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {

	// faqs Routes
	app.Post("/faqs", faq.Store)
	app.Get("/faqs", faq.Index)
	app.Get("/faqs/:id", faq.Show)
	app.Delete("/faqs/:id", faq.Destroy)
	app.Put("/faqs/:id", faq.Update)
	// faqs Routes End

	// User Routes
	app.Post("/users", user.Store)
	//app.Get("/users", user.Index)
	app.Get("/users/:id", user.Show)
	app.Delete("/users/:id", user.Destroy)
	app.Put("/users/:id", user.Update)
	app.Post("/users/login", user.Login)
	// User Routes End

	// User Routes
	app.Post("/profiles", profile.Store)
	app.Get("/profiles/:id", profile.Show)
	app.Delete("/profiles/:id", profile.Destroy)
	app.Put("/profiles/:id", profile.Update)
	// User Routes End
}
