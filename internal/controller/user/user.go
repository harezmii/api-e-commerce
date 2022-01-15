package user

import (
	"api/internal/entity"
	"api/internal/entity/dto"
	"api/internal/entity/response"
	db2 "api/internal/infraStructure/prismaClient"
	prisma "api/internal/infraStructure/prismaClient"
	"api/internal/secret/hash"
	"api/internal/validate"
	"context"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// TODO
// Index empty array

var client = prisma.Client
var contextt = context.Background()

// ShowAccount godoc
// @Summary      Create Data
// @Description  create users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body body  entity.User  false   "User form"
// @Success      201  {object}  []entity.User
// @Router       /users [post]
func Store(ctx *fiber.Ctx) error {
	prisma.PrismaConnection()
	var user entity.User

	parseError := ctx.BodyParser(&user)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}
	err := validate.ValidateStructToTurkish(&user)
	if err == nil {
		passwordHash, passwordHashError := hash.PasswordHash(user.Password)
		if passwordHashError != nil {
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "User not created"})
		}

		createdUser, err := client.User.CreateOne(db2.User.Name.Set(user.Name), db2.User.Surname.Set(user.Surname), db2.User.Email.Set(user.Email), db2.User.Password.Set(passwordHash), db2.User.Status.Set(*user.Status)).Exec(contextt)
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
// @Param        body body  entity.User  false   "User update form"
// @Success      200  {object}  entity.User
// @Router       /users/{id} [put]
func Update(ctx *fiber.Ctx) error {
	prisma.PrismaConnection()
	var user entity.User

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&user)

	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad request"})
	}

	updatedUser, err := client.User.FindUnique(db2.User.ID.Equals(idInt)).Update(db2.User.Name.Set(user.Name), db2.User.Surname.Set(user.Surname), db2.User.Email.Set(user.Email), db2.User.Password.Set(user.Password), db2.User.Status.Set(*user.Status)).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User not updated"})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User updated", Data: updatedUser})

}

//// ShowAccount godoc
//// @Summary      All  Data
//// @Description  get all users
//// @Tags         Users
//// @Accept       json
//// @Produce      json
//// @Param        offset  query  string  true   "Offset"
//// @Success      200  {object}  entity.User
//// @Router       /users [get]
//func Index(ctx *fiber.Ctx) error {
//	var offsetInt int
//	offset := ctx.Query("offset")
//	if offset == "" {
//		offsetInt = 0
//	} else {
//		offsetConvert, convertError := strconv.Atoi(offset)
//		if convertError != nil {
//			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
//		}
//		if offsetConvert >= 0 {
//			offsetInt = offsetConvert
//		} else {
//			offsetInt = 0
//		}
//
//	}
//	prisma.PrismaConnection()
//	allUser, err := client.User.FindMany().Take(10).Skip(offsetInt).Exec(contextt)
//	if err != nil {
//		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User is empty"})
//	}
//	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User is all", Data: allUser})
//}

// ShowAccount godoc
// @Summary      Delete Data
// @Description  delete users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "User ID"
// @Success      200  {object}  []entity.User
// @Router       /users/{id} [delete]
func Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	prisma.PrismaConnection()
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
// @Success      200  {object}  entity.User
// @Router       /users/{id} [get]
func Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	prisma.PrismaConnection()
	singleUser, err := client.User.FindFirst(db2.User.ID.Equals(idInt)).Exec(contextt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User is finding", Data: singleUser})
}

func Login(ctx *fiber.Ctx) error {
	prisma.PrismaConnection()
	var login entity.Login

	parseError := ctx.BodyParser(&login)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}
	errr := validate.ValidateStructToTurkish(login)
	if errr == nil {
		singleUser, err := client.User.FindFirst(db2.User.Email.Equals(login.Email)).Exec(contextt)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
		}
		isLoginSuccess := hash.PasswordHashCompare(login.Password, singleUser.Password)
		if isLoginSuccess {

			return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Login success", Data: dto.LoginUserDTO{UserID: singleUser.ID, Name: singleUser.Name, Surname: singleUser.Surname, Email: singleUser.Email, Status: &singleUser.Status}})
		}

	}
	return nil
}
