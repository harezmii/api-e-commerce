package controller

import (
	"api/ent"
	"context"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Client  *ent.Client
	Context context.Context
	Entity  interface{}
}

type IControl interface {
	Store(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Destroy(ctx *fiber.Ctx) error
	Index(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
