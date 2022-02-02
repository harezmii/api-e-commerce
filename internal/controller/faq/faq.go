package faq

import (
	"api/internal/controller"
	"api/internal/entity"
	"api/internal/entity/dto"
	"api/internal/entity/response"
	"api/internal/logs"
	"api/internal/validate"
	"entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
	_ "net/http"
	"strconv"
	"strings"
	"time"
)

type ControllerFaq struct {
	controller.Controller
}

// Store ShowAccount godoc
// @Summary      Create Data
// @Description  create faqs
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        body body  entity.Faq  false   "Faq form"
// @Success      201  {object}  []entity.Faq
// @Router       /faqs [post]
func (f ControllerFaq) Store(ctx *fiber.Ctx) error {
	faq := f.Entity.(entity.Faq)

	parseError := ctx.BodyParser(&faq)
	if parseError != nil {
		logs.Logger(ctx, "Store!Bad Request , parse error.", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.ErrorResponse{StatusCode: 400, Message: "Bad request , body parse error." + parseError.Error()})
	}
	err := validate.ValidateStructToTurkish(&faq)
	if err == nil {
		dbError := f.Client.Faq.Create().SetQuestion(faq.Question).SetAnswer(faq.Answer).SetStatus(*faq.Status).Exec(f.Context)

		if dbError != nil {
			logs.Logger(ctx, "Store!Faq not created.Database error.", logs.ERROR)
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Faq not created.Database error."})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Faq created", Data: faq})
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: err})
}

// Update ShowAccount godoc
// @Summary      Update Data
// @Description  update faq
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "Faq Id"
// @Param        body body  entity.Faq  false   "Faq update form"
// @Success      200  {object}  entity.Faq
// @Router       /faqs/{id} [put]
func (f ControllerFaq) Update(ctx *fiber.Ctx) error {
	faq := f.Entity.(entity.Faq)

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		logs.Logger(ctx, "Update!Bad Request , Invalid type error. Type must int"+convertError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&faq)

	if parseError != nil {
		logs.Logger(ctx, "Update!Bad Request , parse error."+parseError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}

	validateError := validate.ValidateStructToTurkish(&faq)
	if validateError == nil {
		// Not delete record finding
		selectId, err := f.Client.Faq.Query().Where(func(s *sql.Selector) {
			s.Where(sql.IsNull("deleted_at"))
			s.Where(sql.EQ("id", idInt))
		}).FirstID(f.Context)

		// Not deleting record
		if selectId != 0 {
			errt := f.Client.Faq.UpdateOneID(idInt).SetQuestion(faq.Question).SetAnswer(faq.Answer).SetStatus(*faq.Status).SetUpdatedAt(time.Now()).Exec(f.Context)
			if errt != nil {
				return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "faq not updated, " + strings.Split(errt.Error(), ":")[3]})
			}
		}
		if err != nil {
			logs.Logger(ctx, "Update!Faq not updated.", logs.ERROR)
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "faq not updated"})
		}
		return ctx.Status(fiber.StatusOK).JSON(
			response.SuccessResponse{StatusCode: 200, Message: "Faq Updated.", Data: faq},
		)
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: validateError})
}

// Index ShowAccount godoc
// @Summary      All  Data
// @Description  get all faqs
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        offset  query  string  true   "Offset"
// @Success      200  {object}  entity.Faq
// @Router       /faqs [get]
func (f ControllerFaq) Index(ctx *fiber.Ctx) error {
	var offsetInt int
	offset := ctx.Query("offset")
	if offset == "" {
		offsetInt = 0
	} else {
		offsetConvert, convertError := strconv.Atoi(offset)
		if convertError != nil {
			logs.Logger(ctx, "Index!Bad Request , Invalid type error. Offset type must int", logs.ERROR)
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
		}
		if offsetConvert >= 0 {
			offsetInt = offsetConvert
		} else {
			offsetInt = 0
		}

	}
	var responseDto []dto.FaqDto
	err := f.Client.Faq.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
	}).Limit(10).Offset(offsetInt).Select("id", "question", "answer", "status").Scan(f.Context, &responseDto)
	if err != nil {
		logs.Logger(ctx, "Index!Faq is empty", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Faq is empty"})
	}
	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Faq not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Faq is all", Data: responseDto})
}

// Destroy ShowAccount godoc
// @Summary      Delete Data
// @Description  delete faqs
// @Tags         Faqs
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Faq ID"
// @Success      200  {object}  []entity.Faq
// @Router       /faqs/{id} [delete]
func (f ControllerFaq) Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		//logs.Logger(ctx, "Delete!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	// Not delete record finding
	selectId, err := f.Client.Faq.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).FirstID(f.Context)

	// Not deleting record
	if selectId != 0 {
		f.Client.Faq.UpdateOneID(idInt).SetDeletedAt(time.Now()).Exec(f.Context)
	}
	if err != nil {
		logs.Logger(ctx, "Delete!Faq not find.Not deleted.", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Faq not find.Not deleted."})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Faq deleted", Data: "Faq deleted id:"})
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
func (f ControllerFaq) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	// Id convert error
	if convertError != nil {
		logs.Logger(ctx, "Show!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	var responseDto []dto.FaqDto
	err := f.Client.Faq.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).Select("id", "question", "answer", "status").Scan(f.Context, &responseDto)

	// Database query error
	if err != nil {
		logs.Logger(ctx, "Show!Faq not finding", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Faq not finding"})
	}

	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Faq not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Faq is finding", Data: responseDto})
}
