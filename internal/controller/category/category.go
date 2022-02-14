package category

import (
	"api/ent"
	"api/internal/controller"
	"api/internal/entity"
	"api/internal/entity/dto"
	"api/internal/entity/response"
	"api/internal/logs"
	"api/internal/validate"
	"entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

type ControllerCategory struct {
	controller.Controller
}

// Store ShowAccount godoc
// @Summary      Create Data
// @Description  create categories
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        body body  entity.Category  false   "Category form"
// @Success      201  {object}  []entity.Category
// @Router       /categories [post]
func (c ControllerCategory) Store(ctx *fiber.Ctx) error {
	category := c.Entity.(entity.Category)

	parseError := ctx.BodyParser(&category)
	if parseError != nil {
		logs.Logger(ctx, "Store!Bad Request , parse error.", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.ErrorResponse{StatusCode: 400, Message: "Bad request , body parse error." + parseError.Error()})
	}
	err := validate.ValidateStructToTurkish(&category)
	if err == nil {
		dbError := c.Client.Category.Create().SetImage(category.Image).SetDescription(category.Description).SetDescription(category.Description).SetTitle(category.Title).SetStatus(*category.Status).Exec(c.Context)

		if dbError != nil {
			logs.Logger(ctx, "Store!Category not created.Database error.", logs.ERROR)
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Category not created.Database error."})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Category created", Data: category})
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: err})
}

// Update ShowAccount godoc
// @Summary      Update Data
// @Description  update category
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "Category ID"
// @Param        body body  entity.Category  false   "Category update form"
// @Success      200  {object}  entity.Category
// @Router       /categories/{id} [put]
func (c ControllerCategory) Update(ctx *fiber.Ctx) error {
	category := c.Entity.(entity.Category)

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		logs.Logger(ctx, "Update!Bad Request , Invalid type error. Type must int"+convertError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&category)

	if parseError != nil {
		logs.Logger(ctx, "Update!Bad Request , parse error."+parseError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}

	validateError := validate.ValidateStructToTurkish(&category)
	if validateError == nil {
		// Not delete record finding
		selectId, err := c.Client.Category.Query().Where(func(s *sql.Selector) {
			s.Where(sql.IsNull("deleted_at"))
			s.Where(sql.EQ("id", idInt))
		}).FirstID(c.Context)

		// Not deleting record
		if selectId != 0 {
			errt := c.Client.Category.UpdateOneID(idInt).SetImage(category.Image).SetDescription(category.Description).SetDescription(category.Description).SetTitle(category.Title).SetStatus(*category.Status).SetUpdatedAt(time.Now()).Exec(c.Context)
			if errt != nil {
				return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "category not updated, " + strings.Split(errt.Error(), ":")[3]})
			}
		}
		if err != nil {
			logs.Logger(ctx, "Update!Category not updated.", logs.ERROR)
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "category not updated"})
		}
		return ctx.Status(fiber.StatusOK).JSON(
			response.SuccessResponse{StatusCode: 200, Message: "Faq Updated.", Data: category},
		)
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: validateError})
}

// TODO
/*
	- selectFields için hatalı alan girildiğinde hata dön
*/

// Index ShowAccount godoc
// @Summary      All  Data
// @Description  Get all categories
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        offset  query  string  true   "Offset"
// @Success      200  {object}  entity.Category
// @Router       /categories [get]
func (c ControllerCategory) Index(ctx *fiber.Ctx) error {
	arg := controller.QueryArg{}
	queryParseError := ctx.QueryParser(&arg)

	if queryParseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Query parse error"})
	}
	// Sort Field
	var sortField string
	if arg.SortField == "" {
		sortField = "created_at"
	} else {
		sortField = arg.SortField
	}
	// SortField END

	// SortControl
	var sort ent.OrderFunc
	if arg.Sort == "desc" {
		sort = ent.Desc(sortField)
	} else {
		sort = ent.Asc(sortField)
	}
	// Sort Control END

	// Search Field Control
	var selectField []string
	if arg.SelectFields == "" {
		selectField = []string{"id", "title", "keywords", "description", "image", "status"}
	} else {
		selectField = strings.Split(arg.SelectFields, ",")
	}
	// Search Field Control End

	// Offset Control
	var offsetInt int
	offset := arg.Offset
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
	// Offset Control END

	var responseDto []dto.CategoryDto
	err := c.Client.Category.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
	}).Limit(10).Offset(offsetInt).Order(sort).Select(selectField...).Scan(c.Context, &responseDto)

	if err != nil {
		logs.Logger(ctx, "Index!Faq is empty", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Category is empty"})
	}
	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Category not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Category is all", Data: responseDto})
}

// Destroy ShowAccount godoc
// @Summary      Delete Data
// @Description  delete categories
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Category ID"
// @Success      200  {object}  []entity.Category
// @Router       /categories/{id} [delete]
func (c ControllerCategory) Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		//logs.Logger(ctx, "Delete!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	// Not delete record finding
	selectId, err := c.Client.Category.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).FirstID(c.Context)

	// Not deleting record
	if selectId != 0 {
		c.Client.Faq.UpdateOneID(idInt).SetDeletedAt(time.Now()).Exec(c.Context)
	}
	if err != nil {
		logs.Logger(ctx, "Delete!Category not find.Not deleted.", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Category not find.Not deleted."})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Category deleted", Data: "Category deleted id:"})
}

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Category ID"
// @Success      200  {object}  entity.Category
// @Router       /categories/{id} [get]
func (c ControllerCategory) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	// Id convert error
	if convertError != nil {
		logs.Logger(ctx, "Show!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	var responseDto []dto.CategoryDto
	err := c.Client.Faq.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).Select("id", "title", "keywords", "description", "image", "status").Scan(c.Context, &responseDto)

	// Database query error
	if err != nil {
		logs.Logger(ctx, "Show!Category not finding", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Category not finding"})
	}

	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Category not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Category is finding", Data: responseDto})
}
