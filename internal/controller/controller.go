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

func (Controller) IsOkImageType(contentType string) bool {
	okImageType := []string{"image/png", "image/jpeg", "image/jpg"}
	for _, imageType := range okImageType {
		if imageType == contentType {
			return true
		}
	}
	return false
}

type Control interface {
	Store(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Destroy(ctx *fiber.Ctx) error
	Index(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
