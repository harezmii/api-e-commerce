package user

import (
	"api/internal/controller"
	"api/internal/entity"
	"api/internal/entity/dto"
	"api/internal/entity/response"
	minioUpload "api/internal/infraStructure/minio"
	"api/internal/logs"
	"api/internal/secret/hash"
	"api/internal/secret/jwtManage"
	"api/internal/validate"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid2 "github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

// TODO
// User profile eklenmesini kontrol edeceğim.

type ControllerUser struct {
	controller.Controller
}

// Store ShowAccount godoc
// @Summary      Create Data
// @Description  create users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body body  entity.User  false   "User form"
// @Success      201  {object}  []entity.User
// @Router       /users [post]
func (u ControllerUser) Store(ctx *fiber.Ctx) error {

	user := u.Entity.(entity.User)

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

		dbErr := u.Client.User.Create().SetName(user.Name).SetPassword(passwordHash).SetSurname(user.Surname).SetEmail(user.Email).SetStatus(*user.Status).Exec(u.Context)
		if dbErr != nil {
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "User not created"})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "User created.", Data: user})
	}
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.SuccessResponse{StatusCode: 422, Message: err})

}

// Update ShowAccount godoc
// @Summary      User Update Data
// @Description  update user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "User Id"
// @Param        body body  entity.User  false   "User update form"
// @Success      200  {object}  entity.User
// @Router       /users/{id} [put]
func (u ControllerUser) Update(ctx *fiber.Ctx) error {
	user := u.Entity.(entity.User)

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&user)

	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad request"})
	}
	validateError := validate.ValidateStructToTurkish(&user)
	if validateError == nil {
		selectId, err := u.Client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sql.IsNull("deleted_at"))
			s.Where(sql.EQ("id", idInt))
		}).FirstID(u.Context)
		if selectId != 0 {
			errt := u.Client.User.UpdateOneID(idInt).SetName(user.Name).SetPassword(user.Password).SetEmail(user.Email).SetSurname(user.Surname).SetStatus(*user.Status).SetUpdatedAt(time.Now()).Exec(u.Context)
			if errt != nil {
				return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "user not updated, " + strings.Split(errt.Error(), ":")[3]})
			}
		}

		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User not updated"})
		}

		return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User updated", Data: user})
	}

	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: validateError})
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
//
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
func (u ControllerUser) Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	// Not delete record finding
	selectId, err := u.Client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).FirstID(u.Context)

	// Not deleting record
	if selectId != 0 {
		u.Client.User.UpdateOneID(idInt).SetDeletedAt(time.Now()).Exec(u.Context)
	}
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User not deleted"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User deleted", Data: "User deleted id:"})
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
func (u ControllerUser) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	var responseDto []dto.UserDto
	err := u.Client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).Select("id", "name", "surname", "email", "status").Scan(u.Context, &responseDto)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User not finding"})
	}
	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "User not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "User is finding", Data: responseDto})
}

// UserToProfile ShowAccount godoc
// @Summary      Profile Created or Updated Data
// @Description  Created or Updated Profile
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "User Id"
// @Param        body body  entity.Profile  false   "Profile update fom"
// @Success      201  {object}  entity.Profile
// @Router       /users/{id}/profiles [post]
func (u ControllerUser) UserToProfile(ctx *fiber.Ctx) error {
	profile := entity.Profile{}

	file, errorImage := ctx.FormFile("image")
	if errorImage != nil {
		return errorImage
	}
	extension := strings.Split(file.Filename, ".")[1]
	uuid, _ := uuid2.NewUUID()
	profile.Image = fmt.Sprintf("%s.%s", uuid, extension)

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}

	parseError := ctx.BodyParser(&profile)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad request ,body parse error."})
	}
	err := validate.ValidateStructToTurkish(&profile)
	if err == nil {
		c := minioUpload.ConfigDefault("profiles", file.Header["Content-Type"][0])
		user, _ := u.Client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sql.EQ("id", idInt))
		}).First(u.Context)
		_, profileError := u.Client.Profile.Query().Where(func(s *sql.Selector) {
			s.Where(sql.EQ("user_profiles", idInt))
			s.Where(sql.EQ("deleted_at", nil))
		}).First(u.Context)
		fmt.Println("Profile error:" + profileError.Error())
		// updating the profile, if any
		if profileError == nil {

			openr, _ := file.Open()

			defer openr.Close()

			putError := c.PutImage(profile.Image+"/"+id, openr, file.Size)
			if putError != nil {
				fmt.Println("Put error: " + putError.Error())
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created.Database error."})
			}

			fileImage, fileError := c.GetImage(profile.Image)
			if fileError != nil {
				fmt.Println("File image: " + fileError.Error())
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Profile Image not created."})
			}
			profile.Url = fileImage.String()
			fmt.Println(profile.Url)
			data, errCreate := u.Client.Profile.Update().SetOwner(user).SetOwnerID(idInt).SetPhone(profile.Phone).SetAddress(profile.Address).SetImage(profile.Image).SetURL(profile.Url).Save(u.Context)
			if errCreate != nil {
				fmt.Println(errCreate.Error())
				return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Profile not created"})
			}
			return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Profile created", Data: data})
		} else { // create profile
			openr, _ := file.Open()

			putError := c.PutImage(profile.Image, openr, file.Size)
			if putError != nil {
				fmt.Println("Put error: " + putError.Error())
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created.Database error."})
			}

			fileImage, fileError := c.GetImage(profile.Image)
			if fileError != nil {
				fmt.Println("File image: " + fileError.Error())
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Profile Image not created."})
			}
			profile.Url = fileImage.String()
			data, errCreate := u.Client.Profile.Create().SetOwner(user).SetOwnerID(idInt).SetPhone(profile.Phone).SetAddress(profile.Address).SetImage(profile.Image).SetURL(profile.Url).Save(u.Context)
			if errCreate != nil {
				fmt.Println(errCreate.Error())
				return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Profile not created"})
			}
			return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Profile created", Data: data})
		}
	}
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: err})
}

func (u ControllerUser) Login(ctx *fiber.Ctx) error {
	var login entity.Login

	parseError := ctx.BodyParser(&login)
	if parseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}
	validateError := validate.ValidateStructToTurkish(&login)
	if validateError != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.SuccessResponse{StatusCode: 422, Message: "Validate error", Data: validateError})
	}

	user, err := u.Client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("email", login.Email))
	}).First(u.Context)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}
	isLoginSuccess := hash.PasswordHashCompare(login.Password, user.Password)
	if isLoginSuccess {
		tokenManage := jwtManage.TokenManage{}.Initialize()
		token := tokenManage.NewToken(ctx, user.ID)
		return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Login success", Data: token})
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(response.SuccessResponse{StatusCode: 401, Message: "Unauthorized"})
}
