package faq

import (
	"api/internal/entity"
	"api/internal/entity/response"
	"api/internal/infraStructure/prismaClient"
	db2 "api/internal/infraStructure/prismaClient"
	"api/internal/logs"
	"api/internal/validate"
	"context"
	"github.com/gofiber/fiber/v2"
	_ "net/http"
	"strconv"
)

var client = prisma.Client
var contextt = context.Background()

// ShowAccount godoc
// @Summary      Create Data
// @Description  create faqs
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        body body  entity.Faq  false   "Faq form"
// @Success      201  {object}  []entity.Faq
// @Router       /faqs [post]
func Store(ctx *fiber.Ctx) error {
	prisma.PrismaConnection()
	var faq entity.Faq

	logs.Logger("Bad Request , parse error.", logs.ERROR, ctx.IP())

	parseError := ctx.BodyParser(&faq)
	if parseError != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}
	err := validate.ValidateStructToTurkish(&faq)
	if err == nil {
		createdFaq, err := client.Faq.CreateOne(db2.Faq.Question.Set(faq.Question), db2.Faq.Answer.Set(faq.Answer), db2.Faq.Status.Set(*faq.Status)).Exec(contextt)
		if err != nil {
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Faq not created."})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Faq created", Data: createdFaq})
	}

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: err})

}

// ShowAccount godoc
// @Summary      Update Data
// @Description  update faq
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "Faq Id"
// @Param        body body  entity.Faq  false   "Faq update form"
// @Success      200  {object}  entity.Faq
// @Router       /faqs/{id} [put]
func Update(ctx *fiber.Ctx) error {
	prisma.PrismaConnection()
	var faq entity.Faq

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&faq)

	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request."})
	}

	createdFaq, err := client.Faq.FindUnique(db2.Faq.ID.Equals(idInt)).Update(db2.Faq.Question.Set(faq.Question), db2.Faq.Answer.Set(faq.Answer), db2.Faq.Status.Set(*faq.Status)).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "faq not updated"})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		response.SuccessResponse{StatusCode: 200, Message: "Faq Updated.", Data: createdFaq},
	)

}

// ShowAccount godoc
// @Summary      All  Data
// @Description  get all faqs
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        offset  query  string  true   "Offset"
// @Success      200  {object}  entity.Faq
// @Router       /faqs [get]
func Index(ctx *fiber.Ctx) error {
	var offsetInt int
	offset := ctx.Query("offset")
	if offset == "" {
		offsetInt = 0
	} else {
		offsetConvert, convertError := strconv.Atoi(offset)
		if convertError != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
		}
		if offsetConvert >= 0 {
			offsetInt = offsetConvert
		} else {
			offsetInt = 0
		}

	}
	prisma.PrismaConnection()
	allFaq, err := client.Faq.FindMany().Take(10).Skip(offsetInt).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Faq is empty"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Faq is all", Data: allFaq})
}

// ShowAccount godoc
// @Summary      Delete Data
// @Description  delete faqs
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Faq ID"
// @Success      200  {object}  []entity.Faq
// @Router       /faqs/{id} [delete]
func Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	prisma.PrismaConnection()
	deletedFaq, err := client.Faq.FindUnique(db2.Faq.ID.Equals(idInt)).Delete().Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Faq not deleted"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Faq deleted", Data: deletedFaq})
}

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Faq ID"
// @Success      200  {object}  entity.Faq
// @Router       /faqs/{id} [get]
func Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	prisma.PrismaConnection()
	singleFaq, err := client.Faq.FindFirst(db2.Faq.ID.Equals(idInt)).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Faq not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Faq is finding", Data: singleFaq})
}
