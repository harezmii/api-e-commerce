package rest

import "github.com/gofiber/fiber/v2"

type ConfigAccept struct {
	AcceptLanguages []string
	AcceptEncodings []string
	AcceptCharset   []string
	Accepts         []string
}

type Accept interface {
	NewAccept(ctx *fiber.Ctx)
}

func (c ConfigAccept) NewAccept(ctx *fiber.Ctx) {
	ctx.Accepts(c.Accepts...)
	ctx.AcceptsCharsets(c.AcceptCharset...)
	ctx.AcceptsLanguages(c.AcceptLanguages...)
	ctx.AcceptsEncodings(c.AcceptEncodings...)
}