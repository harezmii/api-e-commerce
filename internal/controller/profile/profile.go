package profile

import (
	"api/internal/controller/user"
	"api/internal/entity/response"
	db2 "api/internal/infraStructure/database"
	"api/internal/validate"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Profile struct {
	UserId  int       `json:"userId" validate:"required"`
	Image   string    `json:"image" form:"image" validate:"required"`
	Address string    `json:"address" form:"address" validate:"required"`
	Phone   string    `json:"phone" form:"phone" validate:"required"`
	User    user.User `json:"user" form:"user" validate:"required"`
}

var client = db2.Client
var contextt = context.Background()

// ShowAccount godoc
// @Summary      Create Data
// @Description  create Profiles
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        body body  Profile  false   "Profile form"
// @Success      201  {object}  []Profile
// @Router       /Profile [post]
func Store(ctx *fiber.Ctx) error {
	db2.PrismaConnection()
	var profile Profile

	parseError := ctx.BodyParser(&profile)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}
	err := validate.ValidateStructToTurkish(&profile)
	if err == nil {
		createdProfile, err := client.Profile.CreateOne(db2.Profile.Image.Set(profile.Image), db2.Profile.Address.Set(profile.Address), db2.Profile.Phone.Set(profile.Phone), db2.Profile.User.Link(db2.User.ID.Equals(profile.UserId)), db2.Profile.UserID.Set(profile.UserId)).Exec(contextt)
		if err != nil {
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Profile not created"})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Profile created", Data: createdProfile})
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
// @Param        body body  Profile  false   "Profile update fom"
// @Success      200  {object}  Profile
// @Router       /profile/{id} [put]
func Update(ctx *fiber.Ctx) error {
	db2.PrismaConnection()
	var profile Profile

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

	updatedProfile, err := client.Profile.FindUnique(db2.Profile.ID.Equals(idInt)).Update(db2.Profile.Image.Set(profile.Image), db2.Profile.Address.Set(profile.Address), db2.Profile.Phone.Set(profile.Phone)).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile not updated"})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile updated", Data: updatedProfile})

}

// ShowAccount godoc
// @Summary      Delete Data
// @Description  delete Profiles
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Profile ID"
// @Success      200  {object}  []Profile
// @Router       /profile/{id} [delete]
func Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	db2.PrismaConnection()
	deletedProfile, err := client.Profile.FindUnique(db2.Profile.ID.Equals(idInt)).Delete().Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile  not deleted"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile deleted", Data: deletedProfile})
}

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Profile ID"
// @Success      200  {object}  Profile
// @Router       /profile/{id} [get]
func Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	db2.PrismaConnection()
	singleProfile, err := client.Profile.FindFirst(db2.Profile.ID.Equals(idInt)).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile  not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile is finding", Data: singleProfile})
}
