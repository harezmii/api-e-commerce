package user

import (
	"context"
	db "e-commerce-api/internal/database"
	"e-commerce-api/internal/validate"
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

var client = db.Client
var contextt = context.Background()

// ShowAccount godoc
// @Summary      Create Data
// @Description  create users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body body  User  false   "User form"
// @Success      200  {object}  []User
// @Router       /user [post]
func Store(ctx *fiber.Ctx) error {
	db.PrismaConnection()
	var user User

	parseError := ctx.BodyParser(&user)
	if parseError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request",
		})
	}
	err := validate.ValidateStructToTurkish(&user)
	if err == nil {
		createdUser, err := client.User.CreateOne(db.User.Name.Set(user.Name), db.User.Surname.Set(user.Surname), db.User.Email.Set(user.Email), db.User.Password.Set(user.Password), db.User.Status.Set(*user.Status)).Exec(contextt)
		if err != nil {
			return ctx.Status(204).JSON(fiber.Map{
				"statusCode": 204,
				"message":    "user is not created",
			})
		}

		return ctx.Status(200).JSON(fiber.Map{
			"statusCode": 200,
			"message":    "user created",
			"data":       createdUser,
		})
	}
	return ctx.JSON(fiber.Map{
		"errors": err,
	})

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
	db.PrismaConnection()
	var user User

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request , Invalid type error. Type must int",
		})
	}
	parseError := ctx.BodyParser(&user)
	fmt.Println(user)
	if parseError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request",
		})
	}

	updatedUser, err := client.User.FindUnique(db.User.ID.Equals(idInt)).Update(db.User.Name.Set(user.Name), db.User.Surname.Set(user.Surname), db.User.Email.Set(user.Email), db.User.Password.Set(user.Password), db.User.Status.Set(*user.Status)).Exec(contextt)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"statusCode": 404,
			"message":    "user is not updated",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"statusCode": 200,
		"message":    "user updated",
		"data":       updatedUser,
	})

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
	db.PrismaConnection()
	allUser, err := client.User.FindMany().Take(10).Skip(offsetInt).Exec(contextt)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"statusCode": 404,
			"message":    "user is empty",
		})
	}
	return ctx.JSON(fiber.Map{
		"statusCode": 200,
		"message":    "user created",
		"data":       allUser,
	})
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
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request , Invalid type error. Type must int",
		})
	}
	db.PrismaConnection()
	deletedUser, err := client.User.FindUnique(db.User.ID.Equals(idInt)).Delete().Exec(contextt)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"statusCode": 404,
			"message":    "user is not deleted",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"statusCode": 200,
		"message":    "user deleted",
		"data":       deletedUser,
	})
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
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request , Invalid type error. Type must int",
		})
	}
	db.PrismaConnection()
	singleUser, err := client.User.FindFirst(db.User.ID.Equals(idInt)).Exec(contextt)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"statusCode": 404,
			"message":    "user is not finding",
		})
	}
	return ctx.JSON(fiber.Map{
		"statusCode": 200,
		"message":    "user is finding",
		"data":       singleUser,
	})
}
