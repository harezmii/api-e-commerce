package profile

import (
	"context"
	db2 "e-commerce-api/internal/infraStructure/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Profile struct {
	UserId  int    `json:"userId" validate:"required"`
	Image   string `json:"image" form:"image" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required"`
}

var client = db2.Client
var contextt = context.Background()

//// ShowAccount godoc
//// @Summary      Create Data
//// @Description  create Profiles
//// @Tags         Profiles
//// @Accept       json
//// @Produce      json
//// @Param        body body  Profile  false   "Profile form"
//// @Success      200  {object}  []Profile
//// @Router       /Profile [post]
//func Store(ctx *fiber.Ctx) error {
//	db.PrismaConnection()
//	var profile Profile
//
//	parseError := ctx.BodyParser(&profile)
//	if parseError != nil {
//		return ctx.Status(400).JSON(fiber.Map{
//			"statusCode":   400,
//			"errorMessage": "Bad Request",
//		})
//	}
//	err := validate.ValidateStructToTurkish(&profile)
//	if err == nil {
//		createdProfile, err := client.Profile.CreateOne(db.Profile.UserID.Set(profile.UserId), db.Profile.Image.Set(profile.Image), db.Profile.Address.Set(profile.Address), db.Profile.Phone.Set(profile.Phone)).Exec(contextt)
//		if err != nil {
//			return ctx.Status(204).JSON(fiber.Map{
//				"statusCode": 204,
//				"message":    "profile is not created",
//			})
//		}
//
//		return ctx.Status(200).JSON(fiber.Map{
//			"statusCode": 200,
//			"message":    "profile created",
//			"data":       createdProfile,
//		})
//	}
//	return ctx.JSON(fiber.Map{
//		"errors": err,
//	})
//
//}

// ShowAccount godoc
// @Summary      Profile Update Data
// @Description  update Profile
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "Profile Id"
// @Param        body body  Profile  false   "Profile update form"
// @Success      200  {object}  Profile
// @Router       /profile/{id} [put]
func Update(ctx *fiber.Ctx) error {
	db2.PrismaConnection()
	var profile Profile

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request , Invalid type error. Type must int",
		})
	}
	parseError := ctx.BodyParser(&profile)
	fmt.Println(profile)
	if parseError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request",
		})
	}

	updatedProfile, err := client.Profile.FindUnique(db2.Profile.ID.Equals(idInt)).Update(db2.Profile.Image.Set(profile.Image), db2.Profile.Address.Set(profile.Address), db2.Profile.Phone.Set(profile.Phone)).Exec(contextt)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"statusCode": 404,
			"message":    "Profile is not updated",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"statusCode": 200,
		"message":    "Profile updated",
		"data":       updatedProfile,
	})

}

// ShowAccount godoc
// @Summary      All  Data
// @Description  get all Profiles
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        offset  query  string  true   "Offset"
// @Success      200  {object}  Profile
// @Router       /profile [get]
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
	db2.PrismaConnection()
	allProfile, err := client.Profile.FindMany().Take(10).Skip(offsetInt).Exec(contextt)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"statusCode": 404,
			"message":    "Profile is empty",
		})
	}
	return ctx.JSON(fiber.Map{
		"statusCode": 200,
		"message":    "Profile created",
		"data":       allProfile,
	})
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
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request , Invalid type error. Type must int",
		})
	}
	db2.PrismaConnection()
	deletedProfile, err := client.Profile.FindUnique(db2.Profile.ID.Equals(idInt)).Delete().Exec(contextt)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"statusCode": 404,
			"message":    "Profile is not deleted",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"statusCode": 200,
		"message":    "Profile deleted",
		"data":       deletedProfile,
	})
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
		return ctx.Status(400).JSON(fiber.Map{
			"statusCode":   400,
			"errorMessage": "Bad Request , Invalid type error. Type must int",
		})
	}
	db2.PrismaConnection()
	singleProfile, err := client.Profile.FindFirst(db2.Profile.ID.Equals(idInt)).Exec(contextt)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"statusCode": 404,
			"message":    "Profile is not finding",
		})
	}
	return ctx.JSON(fiber.Map{
		"statusCode": 200,
		"message":    "Profile is finding",
		"data":       singleProfile,
	})
}
