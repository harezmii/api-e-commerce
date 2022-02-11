package handle

import (
	"api/ent"
	"api/internal/controller/faq"
	"api/internal/controller/message"
	"api/internal/controller/profile"
	"api/internal/controller/user"
	"api/internal/entity"
	"context"
	"github.com/gofiber/fiber/v2"
)

var connection = ent.EntConnection()
var backContext = context.Background()

func SetupRoutes(app fiber.Router) {
	f := faq.ControllerFaq{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Faq{}},
	}

	// faqs Routes
	app.Post("/faqs", f.Store)
	app.Get("/faqs", f.Index)
	app.Get("/faqs/:id", f.Show)
	app.Delete("/faqs/:id", f.Destroy)
	app.Put("/faqs/:id", f.Update)
	// faqs Routes End

	u := user.ControllerUser{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.User{}}}

	// User Routes
	app.Post("/users", u.Store)
	app.Get("/users/:id", u.Show)
	app.Delete("/users/:id", u.Destroy)
	app.Put("/users/:id", u.Update)
	app.Post("/users/login", u.Login)
	//// User Routes End

	p := profile.ControllerProfile{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Profile{}}}

	//// Profile Routes
	app.Post("/profiles", p.Store)
	app.Get("/profiles/:id", p.Show)
	app.Delete("/profiles/:id", p.Destroy)
	app.Put("/profiles/:id", p.Update)
	// User Routes End

	m := message.ControllerMessage{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Message{}}}
	// Message Routes
	app.Post("/messages", m.Store)
	app.Get("/messages", m.Index)
	app.Get("/messages/:id", m.Show)
	app.Delete("/messages/:id", m.Destroy)
	app.Put("/messages/:id", m.Update)
	// Message Routes End
}
