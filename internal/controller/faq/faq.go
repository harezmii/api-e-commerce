package faq

import (
	"context"
	db "e-commerce-api/internal/database"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var client = db.Client
var contextt = context.Background()

func Store(ctx *fiber.Ctx) error {
	db.PrismaConnection()

	createdFaq, err := client.Faq.CreateOne(db.Faq.Question.Set("Bu ilk sorudur"), db.Faq.Answer.Set("ilk sorunun cevabıdır"), db.Faq.Status.Set(true)).Exec(contextt)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"statusCode": 204,
			"message":    "faq is not created",
		})
	}

	return ctx.JSON(fiber.Map{
		"statusCode": 200,
		"message":    "faq created",
		"data":       createdFaq,
	})
}
func Index(ctx *fiber.Ctx) error {
	db.PrismaConnection()
	allFaq, err := client.Faq.FindMany().Exec(contextt)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"statusCode": 404,
			"message":    "faq is empty",
		})
	}
	return ctx.JSON(fiber.Map{
		"statusCode": 200,
		"message":    "faq created",
		"data":       allFaq,
	})
}

func Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, _ := strconv.Atoi(id)
	db.PrismaConnection()
	deletedFaq, err := client.Faq.FindUnique(db.Faq.ID.Equals(idInt)).Delete().Exec(contextt)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"statusCode": 404,
			"message":    "faq is not deleted",
		})
	}
	return ctx.JSON(fiber.Map{
		"statusCode": 200,
		"message":    "faq deleted",
		"data":       deletedFaq,
	})
}
func Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, _ := strconv.Atoi(id)
	db.PrismaConnection()
	singleFaq, err := client.Faq.FindFirst(db.Faq.ID.Equals(idInt)).Exec(contextt)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"statusCode": 404,
			"message":    "faq is not finding",
		})
	}
	return ctx.JSON(fiber.Map{
		"statusCode": 200,
		"message":    "faq is finding",
		"data":       singleFaq,
	})
}
