package faq

import (
	"context"
	db2 "e-commerce-api/internal/infraStructure/database"
	"e-commerce-api/internal/validate"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "net/http"
	"strconv"
)

type Faq struct {
	Question string `json:"question" form:"question" validate:"required"`
	Answer   string `json:"answer" form:"answer" validate:"required"`
	Status   *bool  `json:"status" form:"status" validate:"required"`
}

var client = db2.Client
var contextt = context.Background()

// ShowAccount godoc
// @Summary      Create Data
// @Description  create faqs
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        body body  Faq  false   "Faq form"
// @Success      200  {object}  []Faq
// @Router       /faq [post]
func Store(ctx *fiber.Ctx) error {
	db2.PrismaConnection()
	var faq Faq

	parseError := ctx.BodyParser(&faq)
	fmt.Println(faq)
	if parseError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request",
		})
	}
	err := validate.ValidateStructToTurkish(&faq)
	if err == nil {
		createdFaq, err := client.Faq.CreateOne(db2.Faq.Question.Set(faq.Question), db2.Faq.Answer.Set(faq.Answer), db2.Faq.Status.Set(*faq.Status)).Exec(contextt)
		if err != nil {
			return ctx.Status(204).JSON(fiber.Map{
				"statusCode": 204,
				"message":    "faq is not created",
			})
		}

		return ctx.Status(200).JSON(fiber.Map{
			"statusCode": 200,
			"message":    "faq created",
			"data":       createdFaq,
		})
	}
	return ctx.JSON(fiber.Map{
		"errors": err,
	})

}

// ShowAccount godoc
// @Summary      Update Data
// @Description  update faq
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "Faq Id"
// @Param        body body  Faq  false   "Faq update form"
// @Success      200  {object}  Faq
// @Router       /faq/{id} [put]
func Update(ctx *fiber.Ctx) error {
	db2.PrismaConnection()
	var faq Faq

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request , Invalid type error. Type must int",
		})
	}
	parseError := ctx.BodyParser(&faq)
	fmt.Println(faq)
	if parseError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request",
		})
	}

	createdFaq, err := client.Faq.FindUnique(db2.Faq.ID.Equals(idInt)).Update(db2.Faq.Question.Set(faq.Question), db2.Faq.Answer.Set(faq.Answer), db2.Faq.Status.Set(*faq.Status)).Exec(contextt)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"statusCode": 404,
			"message":    "faq is not updated",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"statusCode": 200,
		"message":    "faq updated",
		"data":       createdFaq,
	})

}

// ShowAccount godoc
// @Summary      All  Data
// @Description  get all faqs
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        offset  query  string  true   "Offset"
// @Success      200  {object}  Faq
// @Router       /faq [get]
func Index(ctx *fiber.Ctx) error {
	var offsetInt int
	offset := ctx.Query("offset")
	if offset == "" {
		offsetInt = 0
	} else {
		offsetConvert, convertError := strconv.Atoi(offset)
		if convertError != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"statusCode":   400,
				"errorMessage": "Bad Request , Invalid type error. Type must int",
			})
		}
		if offsetConvert >= 0 {
			offsetInt = offsetConvert
		} else {
			offsetInt = 0
		}

	}
	db2.PrismaConnection()
	allFaq, err := client.Faq.FindMany().Take(10).Skip(offsetInt).Exec(contextt)
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

// ShowAccount godoc
// @Summary      Delete Data
// @Description  delete faqs
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Faq ID"
// @Success      200  {object}  []Faq
// @Router       /faq/{id} [delete]
func Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request , Invalid type error. Type must int",
		})
	}
	db2.PrismaConnection()
	deletedFaq, err := client.Faq.FindUnique(db2.Faq.ID.Equals(idInt)).Delete().Exec(contextt)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"statusCode": 404,
			"message":    "faq is not deleted",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"statusCode": 200,
		"message":    "faq deleted",
		"data":       deletedFaq,
	})
}

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Faq ID"
// @Success      200  {object}  Faq
// @Router       /faq/{id} [get]
func Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request , Invalid type error. Type must int",
		})
	}
	db2.PrismaConnection()
	singleFaq, err := client.Faq.FindFirst(db2.Faq.ID.Equals(idInt)).Exec(contextt)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
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
