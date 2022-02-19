package handle

import (
	"api/ent"
	"api/internal/controller/category"
	"api/internal/controller/faq"
	"api/internal/controller/image"
	"api/internal/controller/message"
	"api/internal/controller/profile"
	"api/internal/controller/user"
	"api/internal/entity"
	"context"
	"github.com/gofiber/fiber/v2"
)

var connection = ent.ConnectionEnt()
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
	app.Post("/users/:id/profiles", u.UserToProfile)
	// User Routes End

	p := profile.ControllerProfile{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Profile{}}}

	//// Profile Routes
	app.Get("/profiles/:id", p.Show)
	app.Delete("/profiles/:id", p.Destroy)
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

	c := category.ControllerCategory{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Category{}},
	}
	app.Post("/categories", c.Store)
	app.Get("/categories", c.Index)
	app.Get("/categories/:id", c.Show)
	app.Delete("/categories/:id", c.Destroy)
	app.Put("/categories/:id", c.Update)

	co := category.ControllerCategory{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Comment{}},
	}
	app.Post("/comments", co.Store)
	app.Get("/comments", co.Index)
	app.Get("/comments/:id", co.Show)
	app.Delete("/comments/:id", co.Destroy)
	app.Put("/comments/:id", co.Update)

	i := image.ControllerImage{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Image{}},
	}
	app.Post("/images", i.Store)
	app.Get("/images", i.Index)
	app.Get("/images/:id", i.Show)
	app.Delete("/images/:id", i.Destroy)
	app.Put("/images/:id", i.Update)
}
