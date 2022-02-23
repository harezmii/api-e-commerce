package image

import (
	"api/ent"
	"api/internal/controller"
	"api/internal/entity"
	"api/internal/entity/dto"
	"api/internal/entity/response"
	"api/internal/infraStructure/minio"
	"api/internal/logs"
	"api/internal/validate"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid2 "github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

type ControllerImage struct {
	controller.Controller
}

//Store ShowAccount godoc
//@Summary      Create Data
//@Description  create Image
//@Tags         Images
//@Accept       json
//@Produce      json
//@Param        body body  entity.Image  false   "Image form"
//@Success      201  {object}  []entity.Image
//@Router       /images [post]
func (i ControllerImage) Store(ctx *fiber.Ctx) error {
	file, errorImage := ctx.FormFile("image")
	if errorImage != nil {
		return errorImage
	}

	extension := strings.Split(file.Filename, ".")[1]
	uuid, _ := uuid2.NewUUID()
	image := i.Entity.(entity.Image)
	image.Image = fmt.Sprintf("%s.%s", uuid, extension)

	parseError := ctx.BodyParser(&image)

	if parseError != nil {
		logs.Logger(ctx, "Store!Bad Request , parse error.", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.ErrorResponse{StatusCode: 400, Message: "Bad request , body parse error." + parseError.Error()})
	}
	err := validate.ValidateStructToTurkish(&image)
	if err == nil {

		openr, _ := file.Open()

		defer openr.Close()

		c := minioUpload.ConfigDefault("image", file.Header["Content-Type"][0])
		putError := c.PutImage(image.Image, openr, file.Size)
		if putError != nil {
			fmt.Println("Put error: " + putError.Error())
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created.Database error."})
		}

		fileImage, fileError := c.GetImage(image.Image)
		if fileError != nil {
			fmt.Println("File image: " + fileError.Error())
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created."})
		}
		image.Url = fileImage.String()

		dbError := i.Client.Image.Create().SetImage(image.Image).SetTitle(image.Title).SetURL(image.Url).SetStatus(true).Exec(i.Context)
		if dbError != nil {
			fmt.Println("Db Error: " + dbError.Error())
			logs.Logger(ctx, "Store!Image not created.Database error.", logs.ERROR)
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created.Database error."})
		}
		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Image created", Data: image})
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: err})
}

// Update ShowAccount godoc
// @Summary      Update Data
// @Description  update images
// @Tags         Images
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "images ID"
// @Param        body body  entity.Image  false   "image update form"
// @Success      200  {object}  entity.Image
// @Router       /images/{id} [put]
func (i ControllerImage) Update(ctx *fiber.Ctx) error {
	file, errorImage := ctx.FormFile("image")
	if errorImage != nil {
		return errorImage
	}

	image := i.Entity.(entity.Image)

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		logs.Logger(ctx, "Update!Bad Request , Invalid type error. Type must int"+convertError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&image)

	if parseError != nil {
		logs.Logger(ctx, "Update!Bad Request , parse error."+parseError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}

	validateError := validate.ValidateStructToTurkish(&image)
	if validateError == nil {
		// Not delete record finding
		selectId, err := i.Client.Image.Query().Where(func(s *sql.Selector) {
			s.Where(sql.IsNull("deleted_at"))
			s.Where(sql.EQ("id", idInt))
		}).First(i.Context)

		// Not deleting record
		if selectId.ID != 0 {
			open, _ := file.Open()
			defer open.Close()
			cfg := minioUpload.ConfigDefault("image", "")

			putError := cfg.PutImage(selectId.Image, open, file.Size)
			if putError != nil {
				fmt.Println(putError.Error())
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not updated"})
			}
			getImage, getError := cfg.GetImage(selectId.Image)
			if getError != nil {
				fmt.Println(getError.Error())
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not updated"})

			}
			image.Url = getImage.String()
			errt := i.Client.Image.UpdateOneID(idInt).SetImage(selectId.Image).SetTitle(image.Title).SetURL(getImage.String()).Exec(i.Context)
			if errt != nil {
				return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image not updated, " + strings.Split(errt.Error(), ":")[3]})
			}
		}
		if err != nil {
			logs.Logger(ctx, "Update!Image not updated.", logs.ERROR)
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image not updated"})
		}
		image.Image = selectId.Image
		return ctx.Status(fiber.StatusOK).JSON(
			response.SuccessResponse{StatusCode: 200, Message: "Image Updated.", Data: image},
		)
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: validateError})
}

// Index ShowAccount godoc
// @Summary      All  Data
// @Description  Get all images
// @Tags         Images
// @Accept       json
// @Produce      json
// @Param        offset  query  string  true   "Offset"
// @Success      200  {object}  entity.Image
// @Router       /images [get]
func (i ControllerImage) Index(ctx *fiber.Ctx) error {
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
		selectField = []string{"id", "title", "image"}
	} else {
		selectField = strings.Split(arg.SelectFields, ",")
	}
	// Search Field Control End

	// Offset Control
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
	// Offset Control END

	var responseDto []dto.ImageDto
	err := i.Client.Image.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
	}).Limit(10).Offset(offsetInt).Order(sort).Select(selectField...).Scan(i.Context, &responseDto)

	if err != nil {
		logs.Logger(ctx, "Index!Image is empty", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image is empty"})
	}
	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image not finding"})
	}

	for i := 0; i < len(responseDto); i++ {
		cfg := minioUpload.ConfigDefault("image", "")
		file, fileError := cfg.GetImage(responseDto[i].Image)
		if fileError != nil {
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created."})
		}
		responseDto[i].Url = file.String()
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Image is all", Data: responseDto})
}

// Destroy ShowAccount godoc
// @Summary      Delete Data
// @Description  delete images
// @Tags         Images
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Image ID"
// @Success      200  {object}  []entity.Image
// @Router       /images/{id} [delete]
func (i ControllerImage) Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		//logs.Logger(ctx, "Delete!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	// Not delete record finding
	firstImage, err := i.Client.Image.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).First(i.Context)
	if firstImage == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image not find.Not deleted."})
	}
	// Not deleting record
	if firstImage.ID != 0 {
		i.Client.Image.UpdateOneID(idInt).SetDeletedAt(time.Now()).Exec(i.Context)
	}
	if err != nil {
		logs.Logger(ctx, "Delete!Image not find.Not deleted.", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image not find.Not deleted."})
	}

	// MINIO Remove Object
	cfg := minioUpload.ConfigDefault("image", "")
	cfg.RemoveImage(firstImage.Image)
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Image deleted", Data: "Image deleted id:"})
}

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Images
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Image ID"
// @Success      200  {object}  entity.Image
// @Router       /images/{id} [get]
func (i ControllerImage) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	// Id convert error
	if convertError != nil {
		logs.Logger(ctx, "Show!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	var responseDto []dto.ImageDto
	err := i.Client.Image.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).Select("id", "title", "image", "url").Scan(i.Context, &responseDto)
	// Database query error
	if err != nil {
		//logs.Logger(ctx, "Show!Image not finding", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image not finding"})
	}

	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image not finding"})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Image is finding", Data: responseDto})
}
