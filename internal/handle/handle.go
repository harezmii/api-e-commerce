package handle

import (
	"api/ent"
	"api/internal/controller/category"
	"api/internal/controller/comment"
	"api/internal/controller/faq"
	"api/internal/controller/message"
	"api/internal/controller/product"
	"api/internal/controller/profile"
	"api/internal/controller/user"
	"api/internal/entity"
	"context"
	"github.com/gofiber/fiber/v2"
)

var connection = ent.ConnectionEnt()
var backContext = context.Background()

func faqRoutes(app fiber.Router) {
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
}

func userRoutes(app fiber.Router) {
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
	app.Get("/users/:id/products", u.UserOwnerProduct)
	app.Post("/users/login", u.Login)
	// User Routes End
}

func profileRoutes(app fiber.Router) {
	p := profile.ControllerProfile{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Profile{}}}

	//// Profile Routes
	app.Get("/users/:id/profiles", p.Show)
	app.Post("/users/:id/profiles", p.StoreOrUpdate)
	app.Delete("/users/:id/profiles", p.Destroy)
	// Profile Routes End
}

func messageRoutes(app fiber.Router) {
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

func categoryRoutes(app fiber.Router) {
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
	app.Get("/categories/:id/products", c.CategoryOwnProducts)
	app.Delete("/categories/:id", c.Destroy)
	app.Put("/categories/:id", c.Update)
}

func commentRoutes(app fiber.Router) {
	co := comment.ControllerComment{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Comment{}},
	}
	app.Post("/comments", co.Store)
	app.Get("/comments", co.Index)
	app.Get("/comments/:id", co.Show)
	app.Get("/comments/:id/users", co.CommentOwnUsers)
	app.Delete("/comments/:id", co.Destroy)
	app.Put("/comments/:id", co.Update)
}

func productRoutes(app fiber.Router) {
	pr := product.ControllerProduct{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Entity  interface{}
		}{Client: connection, Context: backContext, Entity: entity.Product{}},
	}
	app.Post("/products", pr.Store)
	app.Delete("/products/:id", pr.Destroy)
	app.Get("/products/:id", pr.Show)
	app.Get("/products/:id/comments", pr.CommentOwnProducts)
	app.Put("/products/:id/images/:imageId", pr.Update)
	app.Get("/products", pr.Index)
}
func SetupRoutes(app fiber.Router) {
	faqRoutes(app)
	userRoutes(app)
	profileRoutes(app)
	messageRoutes(app)
	categoryRoutes(app)
	commentRoutes(app)
	productRoutes(app)
}
