package user

import (
	"api/internal/entity/response"
	db2 "api/internal/infraStructure/database"
	"api/internal/validate"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// TODO
// Index empty array

type User struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Surname  string `json:"surname" form:"surname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Status   *bool  `json:"status" form:"status" validate:"required"`
}

var client = db2.Client
var contextt = context.Background()

// ShowAccount godoc
// @Summary      Create Data
// @Description  create users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body body  User  false   "User form"
// @Success      201  {object}  []User
// @Router       /user [post]
func Store(ctx *fiber.Ctx) error {
	db2.PrismaConnection()
	var user User

	parseError := ctx.BodyParser(&user)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}
	err := validate.ValidateStructToTurkish(&user)
	if err == nil {
		createdUser, err := client.User.CreateOne(db2.User.Name.Set(user.Name), db2.User.Surname.Set(user.Surname), db2.User.Email.Set(user.Email), db2.User.Password.Set(user.Password), db2.User.Status.Set(*user.Status)).Exec(contextt)
		if err != nil {
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "User not created"})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "User created.", Data: createdUser})
	}
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.SuccessResponse{StatusCode: 422, Message: err})

}

// ShowAccount godoc
// @Summary      User Update Data
// @Description  update user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "User Id"
// @Param        body body  User  false   "User update form"
// @Success      200  {object}  User
// @Router       /user/{id} [put]
func Update(ctx *fiber.Ctx) error {
	db2.PrismaConnection()
	var user User

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&user)
	fmt.Println(user)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad request"})
	}

	updatedUser, err := client.User.FindUnique(db2.User.ID.Equals(idInt)).Update(db2.User.Name.Set(user.Name), db2.User.Surname.Set(user.Surname), db2.User.Email.Set(user.Email), db2.User.Password.Set(user.Password), db2.User.Status.Set(*user.Status)).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User not updated"})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User updated", Data: updatedUser})

}

// ShowAccount godoc
// @Summary      All  Data
// @Description  get all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        offset  query  string  true   "Offset"
// @Success      200  {object}  User
// @Router       /user [get]
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
	db2.PrismaConnection()
	allUser, err := client.User.FindMany().Take(10).Skip(offsetInt).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User is empty"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User is all", Data: allUser})
}

// ShowAccount godoc
// @Summary      Delete Data
// @Description  delete users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "User ID"
// @Success      200  {object}  []User
// @Router       /user/{id} [delete]
func Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	db2.PrismaConnection()
	deletedUser, err := client.User.FindUnique(db2.User.ID.Equals(idInt)).Delete().Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User not deleted"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User deleted", Data: deletedUser})
}

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "User ID"
// @Success      200  {object}  User
// @Router       /user/{id} [get]
func Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	db2.PrismaConnection()
	singleUser, err := client.User.FindFirst(db2.User.ID.Equals(idInt)).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User is finding", Data: singleUser})
}
