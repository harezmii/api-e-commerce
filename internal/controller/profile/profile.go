package profile

import (
	"api/internal/controller"
	"api/internal/entity"
	"api/internal/entity/dto"
	"api/internal/entity/response"
	"api/internal/logs"
	"api/internal/validate"
	"entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type ControllerProfile struct {
	controller.Controller
}

// Store ShowAccount godoc
// @Summary      Create Data
// @Description  create Profiles
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        body body  entity.Profile  false   "Profile form"
// @Success      201  {object}  []entity.Profile
// @Router       /profiles [post]
func (p ControllerProfile) Store(ctx *fiber.Ctx) error {
	profile := p.Entity.(entity.Profile)

	parseError := ctx.BodyParser(&profile)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad request ,body parse error."})
	}
	err := validate.ValidateStructToTurkish(&profile)
	if err == nil {
		dbError := p.Client.Profile.Create().SetPhone(profile.Phone).SetAddress(profile.Address).SetImage(profile.Image).SetOwnerID(profile.UserId).Exec(p.Context)
		if dbError != nil {
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Profile not created.Database error."})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Profile created.", Data: profile})
	}
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: err})

}

// Update ShowAccount godoc
// @Summary      Profile Update Data
// @Description  update Profile
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "Profile Id"
// @Param        body body  entity.Profile  false   "Profile update fom"
// @Success      200  {object}  entity.Profile
// @Router       /profiles/{id} [put]
func (p ControllerProfile) Update(ctx *fiber.Ctx) error {
	profile := p.Entity.(entity.Profile)

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&profile)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad request"})
	}
	validateError := validate.ValidateStructToTurkish(&profile)
	if validateError == nil {
		// Not delete record finding
		selectId, err := p.Client.Profile.Query().Where(func(s *sql.Selector) {
			s.Where(sql.IsNull("deleted_at"))
			s.Where(sql.EQ("id", idInt))
		}).FirstID(p.Context)
		if selectId != 0 {
			errt := p.Client.Profile.UpdateOneID(idInt).SetPhone(profile.Phone).SetAddress(profile.Address).SetImage(profile.Image).SetUpdatedAt(time.Now()).Exec(p.Context)
			if errt != nil {
				return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile not updated"})
			}
		}
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile not updated"})
		}
		return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile updated", Data: profile})
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: validateError})
}

// Destroy ShowAccount godoc
// @Summary      Delete Data
// @Description  delete Profiles
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Profile ID"
// @Success      200  {object}  []entity.Profile
// @Router       /profiles/{id} [delete]
func (p ControllerProfile) Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	selectId, err := p.Client.Profile.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).FirstID(p.Context)
	if selectId != 0 {
		p.Client.Profile.UpdateOneID(idInt).SetDeletedAt(time.Now()).Exec(p.Context)
	}
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile  not deleted"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile deleted", Data: "Profile deleted id:"})
}

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Profile ID"
// @Success      200  {object}  entity.Profile
// @Router       /profiles/{id} [get]
func (p ControllerProfile) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	var responseDto []dto.ProfileDto
	err := p.Client.Profile.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).Select("id", "address", "phone", "image").Scan(p.Context, &responseDto)

	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile  not finding"})

	}
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile  not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile is finding", Data: responseDto})
}
