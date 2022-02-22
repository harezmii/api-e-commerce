package profile

import (
	"api/internal/controller"
	"api/internal/entity"
	"api/internal/entity/dto"
	"api/internal/entity/response"
	minioUpload "api/internal/infraStructure/minio"
	"api/internal/validate"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid2 "github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

type ControllerProfile struct {
	controller.Controller
}

// StoreOrUpdate UserToProfile ShowAccount godoc
// @Summary      Profile Created or Updated Profile
// @Description  Created or Updated Profile
// @Tags         Profiles
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "User Id"
// @Param        body body  entity.Profile  false   "Profile store or update fom"
// @Success      201  {object}  entity.Profile
// @Router       /users/{id}/profiles [post]
func (p ControllerProfile) StoreOrUpdate(ctx *fiber.Ctx) error {
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
		user, _ := p.Client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sql.EQ("id", idInt))
		}).First(p.Context)

		_, profileError := p.Client.Profile.Query().Where(func(s *sql.Selector) {
			s.Where(sql.EQ("user_profile", user.ID))
			s.Where(sql.EQ("deleted_at", nil))
		}).First(p.Context)

		// updating the profile, if any
		if profileError == nil {

			openr, _ := file.Open()

			defer openr.Close()

			putError := c.PutImage(id+"/"+profile.Image, openr, file.Size)
			if putError != nil {
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created.Database error."})
			}

			fileImage, fileError := c.GetImage(id + "/" + profile.Image)
			if fileError != nil {
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Profile Image not created."})
			}
			profile.Url = fileImage.String()

			_, errCreate := p.Client.Profile.Create().SetOwner(user).SetOwnerID(idInt).SetPhone(profile.Phone).SetAddress(profile.Address).SetImage(profile.Image).SetURL(profile.Url).Save(p.Context)
			if errCreate != nil {
				fmt.Println("Update:" + errCreate.Error())
				return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Profile not created"})
			}
			return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Profile created", Data: profile})
		} else { // create profile

			openr, _ := file.Open()

			putError := c.PutImage(id+"/"+profile.Image, openr, file.Size)
			if putError != nil {
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created.Database error."})
			}

			fileImage, fileError := c.GetImage(id + "/" + profile.Image)
			if fileError != nil {
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Profile Image not created."})
			}
			profile.Url = fileImage.String()

			_, errCreate := p.Client.Profile.Update().SetOwner(user).SetOwnerID(idInt).SetPhone(profile.Phone).SetAddress(profile.Address).SetImage(profile.Image).SetURL(profile.Url).Save(p.Context)
			if errCreate != nil {
				fmt.Println("Create:" + errCreate.Error())
				return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Profile not created"})
			}
			return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Profile created", Data: profile})
		}
	}
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: err})
}

// Destroy ShowAccount godoc
// @Summary      Delete Profile
// @Description  User Delete Profiles
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
		s.Where(sql.EQ("user_profile", idInt))
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
// @Summary      Show Profile
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
		s.Where(sql.EQ("user_profile", idInt))
	}).Select("id", "address", "phone", "image", "url").Scan(p.Context, &responseDto)

	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile  not finding"})

	}
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Profile  not finding"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Profile is finding", Data: responseDto})
}
