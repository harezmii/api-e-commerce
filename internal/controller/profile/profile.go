package profile

import (
	"api/internal/controller"
	"api/internal/entity"
	"api/internal/entity/response"
	"api/internal/validate"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ControllerProfile struct {
	controller.Controller
}

// ShowAccount godoc
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

// ShowAccount godoc
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
	fmt.Println(profile)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad request"})
	}

	err := p.Client.Profile.UpdateOneID(idInt).SetPhone(profile.Phone).SetAddress(profile.Address).SetImage(profile.Image).Exec(p.Context)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile not updated"})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile updated", Data: profile})

}

// ShowAccount godoc
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
	err := p.Client.Profile.DeleteOneID(idInt).Exec(p.Context)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile  not deleted"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile deleted", Data: "deletedProfile"})
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

	singleProfile, err := p.Client.Profile.Get(p.Context, idInt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile  not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile is finding", Data: singleProfile})
}
