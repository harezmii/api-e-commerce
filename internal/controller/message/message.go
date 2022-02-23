package message

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
	_ "net/http"
	"strconv"
	"strings"
	"time"
)

type ControllerMessage struct {
	controller.Controller
}

// Store ShowAccount godoc
// @Summary      Create Data
// @Description  create messages
// @Tags         Messages
// @Accept       json
// @Produce      json
// @Param        body body  entity.Message  false   "Message form"
// @Success      201  {object}  []entity.Message
// @Router       /messages [post]
func (m ControllerMessage) Store(ctx *fiber.Ctx) error {
	message := m.Entity.(entity.Message)

	parseError := ctx.BodyParser(&message)
	if parseError != nil {
		logs.Logger(ctx, "Store!Bad Request , parse error.", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.ErrorResponse{StatusCode: 400, Message: "Bad request , body parse error." + parseError.Error()})
	}
	err := validate.ValidateStructToTurkish(&message)
	if err == nil {
		dbError := m.Client.Message.Create().SetName(message.Name).SetEmail(message.Email).SetIP(message.IP).SetMessage(message.Message).SetStatus(*message.Status).SetPhone(message.Phone).SetSubject(message.Subject).Exec(m.Context)
		if dbError != nil {
			logs.Logger(ctx, "Store!Message not created.Database error.", logs.ERROR)
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Message not created.Database error."})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Message created", Data: message})
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: err})
}

// Update ShowAccount godoc
// @Summary      Update Data
// @Description  update message
// @Tags         Messages
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "Message ID"
// @Param        body body  entity.Message  false   "Message update form"
// @Success      200  {object}  entity.Message
// @Router       /messages/{id} [put]
func (m ControllerMessage) Update(ctx *fiber.Ctx) error {
	message := m.Entity.(entity.Message)

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		logs.Logger(ctx, "Update!Bad Request , Invalid type error. Type must int"+convertError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&message)

	if parseError != nil {
		logs.Logger(ctx, "Update!Bad Request , parse error."+parseError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}
	validateError := validate.ValidateStructToTurkish(&message)
	if validateError == nil {
		// Not delete record finding
		selectId, err := m.Client.Message.Query().Where(func(s *sql.Selector) {
			s.Where(sql.IsNull("deleted_at"))
			s.Where(sql.EQ("id", idInt))
		}).FirstID(m.Context)

		// Not deleting record
		if selectId != 0 {
			errt := m.Client.Message.UpdateOneID(idInt).SetName(message.Name).SetEmail(message.Email).SetIP(message.IP).SetMessage(message.Message).SetStatus(*message.Status).SetPhone(message.Phone).SetSubject(message.Subject).SetUpdatedAt(time.Now()).Exec(m.Context)
			if errt != nil {
				return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Message not updated, " + strings.Split(errt.Error(), ":")[3]})
			}
		}
		if err != nil {
			logs.Logger(ctx, "Update!Message not updated.", logs.ERROR)
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Message not updated"})
		}
		return ctx.Status(fiber.StatusOK).JSON(
			response.SuccessResponse{StatusCode: 200, Message: "Message Updated.", Data: message},
		)
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: validateError})
}

// Index ShowAccount godoc
// @Summary      All  Data
// @Description  Get all messages
// @Tags         Messages
// @Accept       json
// @Produce      json
// @Param        offset  query  string  true   "Offset"
// @Success      200  {object}  entity.Message
// @Router       /messages [get]
func (m ControllerMessage) Index(ctx *fiber.Ctx) error {
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
		selectField = []string{"id", "name", "email", "phone", "subject", "message", "ip"}
	} else {
		selectField = strings.Split(arg.SelectFields, ",")
	}
	// Search Field Control End

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
	var responseDto []dto.MessageDto
	err := m.Client.Message.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
	}).Limit(10).Offset(offsetInt).Order(sort).Select(selectField...).Scan(m.Context, &responseDto)
	if err != nil {
		logs.Logger(ctx, "Index!Message is empty", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Message is empty"})
	}
	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Message not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Message is all", Data: responseDto})
}

// Destroy ShowAccount godoc
// @Summary      Delete Data
// @Description  delete messages
// @Tags         Messages
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Message ID"
// @Success      200  {object}  []entity.Message
// @Router       /messages/{id} [delete]
func (m ControllerMessage) Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		logs.Logger(ctx, "Delete!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}

	// Not delete record finding
	selectId, err := m.Client.Message.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).FirstID(m.Context)

	// Not deleting record
	if selectId != 0 {
		m.Client.Message.UpdateOneID(idInt).SetDeletedAt(time.Now()).Exec(m.Context)
	}
	if err != nil {
		logs.Logger(ctx, "Delete!Message not find.Not deleted.", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Message not find.Not deleted."})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Message deleted", Data: "Message deleted id:"})
}

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Messages
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Message ID"
// @Success      200  {object}  entity.Message
// @Router       /messages/{id} [get]
func (m ControllerMessage) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	// Id convert error
	if convertError != nil {
		logs.Logger(ctx, "Show!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	var responseDto []dto.MessageDto
	err := m.Client.Message.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).Select("id", "name", "email", "phone", "subject", "message", "ip").Scan(m.Context, &responseDto)

	// Database query error
	if err != nil {
		logs.Logger(ctx, "Show!Message not finding", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Message not finding"})
	}

	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Message not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Message is finding", Data: responseDto})
}
