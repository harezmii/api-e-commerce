package product

import (
	"api/ent"
	"api/internal/controller"
	"api/internal/entity"
	"api/internal/entity/dto"
	"api/internal/entity/response"
	minioUpload "api/internal/infraStructure/minio"
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

type ControllerProduct struct {
	controller.Controller
}

// Store ShowAccount godoc
// @Summary      Create Data
// @Description  create products
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        body body  entity.Product  false   "Product form"
// @Success      201  {object}  []entity.Product
// @Router       /products [post]
func (p ControllerProduct) Store(ctx *fiber.Ctx) error {

	file, errorImage := ctx.FormFile("image")
	if errorImage != nil {
		return errorImage
	}

	extension := strings.Split(file.Filename, ".")[1]
	uuid, _ := uuid2.NewUUID()
	product := p.Entity.(entity.Product)
	product.Image = fmt.Sprintf("%s.%s", uuid, extension)
	parseError := ctx.BodyParser(&product)
	if parseError != nil {
		logs.Logger(ctx, "Store!Bad Request , parse error.", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			response.ErrorResponse{StatusCode: 400, Message: "Bad request , body parse error." + parseError.Error()})
	}
	err := validate.ValidateStructToTurkish(&product)
	if err == nil {

		openr, _ := file.Open()

		defer openr.Close()

		c := minioUpload.ConfigDefault("products", file.Header["Content-Type"][0])
		putError := c.PutImage(product.Image, openr, file.Size)
		if putError != nil {
			fmt.Println("Put error: " + putError.Error())
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created.Database error."})
		}

		fileImage, fileError := c.GetImage(product.Image)
		if fileError != nil {
			fmt.Println("File image: " + fileError.Error())
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Image not created."})
		}
		product.Url = fileImage.String()

		dbError := p.Client.Product.Create().SetTitle(product.Title).SetKeywords(product.Keywords).SetDescription(product.Description).SetStatus(*product.Status).SetImage(product.Image).SetURL(product.Url).Exec(p.Context)

		if dbError != nil {
			logs.Logger(ctx, "Store!Product not created.Database error.", logs.ERROR)
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Product not created.Database error."})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse{StatusCode: 201, Message: "Product created", Data: product})
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: err})
}

// Destroy ShowAccount godoc
// @Summary      Delete Data
// @Description  delete products
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "Product ID"
// @Success      200  {object}  []entity.Product
// @Router       /products/{id} [delete]
func (p ControllerProduct) Destroy(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		//logs.Logger(ctx, "Delete!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	// Not delete record finding
	firstImage, err := p.Client.Product.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
	}).First(p.Context)
	if firstImage == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image not find.Not deleted."})
	}
	// Not deleting record
	if firstImage.ID != 0 {
		p.Client.Product.UpdateOneID(idInt).SetDeletedAt(time.Now()).Exec(p.Context)
	}
	if err != nil {
		logs.Logger(ctx, "Delete!Image not find.Not deleted.", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Image not find.Not deleted."})
	}

	// MINIO Remove Object
	cfg := minioUpload.ConfigDefault("products", "")
	cfg.RemoveImage(firstImage.Image)
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Image deleted", Data: "Image deleted id= " + strconv.Itoa(firstImage.ID)})
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
func (p ControllerProduct) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	// Id convert error
	if convertError != nil {
		logs.Logger(ctx, "Show!Bad Request , Invalid type error. Type must int", logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	var responseDto []dto.ProductDto
	err := p.Client.Product.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).Select("id", "title", "keywords", "description", "image", "url").Scan(p.Context, &responseDto)
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

// Index ShowAccount godoc
// @Summary      All  Data
// @Description  Get all products
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        offset  query  string  true   "Offset"
// @Success      200  {object}  entity.Product
// @Router       /products [get]
func (p ControllerProduct) Index(ctx *fiber.Ctx) error {
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
		selectField = []string{"id", "title", "keywords", "description", "image", "url"}
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

	var responseDto []dto.ProductDto
	err := p.Client.Product.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
	}).Limit(10).Offset(offsetInt).Order(sort).Select(selectField...).Scan(p.Context, &responseDto)

	if err != nil {
		logs.Logger(ctx, "Index!Product is empty", logs.ERROR)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Product is empty"})
	}
	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Product not finding"})
	}

	for i := 0; i < len(responseDto); i++ {
		cfg := minioUpload.ConfigDefault("products", "")
		file, fileError := cfg.GetImage(responseDto[i].Image)
		if fileError != nil {
			return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Product not created."})
		}
		responseDto[i].Url = file.String()
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: 200, Message: "Product is all", Data: responseDto})
}

// Update ShowAccount godoc
// @Summary      Update Data
// @Description  update products
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "Product ID"
// @Param        body body  entity.Product  false   "Product update form"
// @Success      200  {object}  entity.Product
// @Router       /products/{id} [put]
func (p ControllerProduct) Update(ctx *fiber.Ctx) error {
	file, errorImage := ctx.FormFile("image")
	if errorImage != nil {
		return errorImage
	}

	extension := strings.Split(file.Filename, ".")[1]
	uuid, _ := uuid2.NewUUID()
	product := p.Entity.(entity.Product)
	product.Image = fmt.Sprintf("%s.%s", uuid, extension)

	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		logs.Logger(ctx, "Update!Bad Request , Invalid type error. Type must int"+convertError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	parseError := ctx.BodyParser(&product)

	if parseError != nil {
		logs.Logger(ctx, "Update!Bad Request , parse error."+parseError.Error(), logs.ERROR)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , parse error."})
	}

	validateError := validate.ValidateStructToTurkish(&product)
	if validateError == nil {
		// Not delete record finding
		selectId, err := p.Client.Product.Query().Where(func(s *sql.Selector) {
			s.Where(sql.IsNull("deleted_at"))
			s.Where(sql.EQ("id", idInt))
		}).First(p.Context)

		// Not deleting record
		if selectId.ID != 0 {
			open, _ := file.Open()
			defer open.Close()
			cfg := minioUpload.ConfigDefault("products", "")

			putError := cfg.PutImage(selectId.Image, open, file.Size)
			if putError != nil {
				fmt.Println(putError.Error())
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Product not updated"})
			}
			getImage, getError := cfg.GetImage(selectId.Image)
			if getError != nil {
				fmt.Println(getError.Error())
				return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: 204, Message: "Product not updated"})

			}
			product.Url = getImage.String()
			errt := p.Client.Product.UpdateOneID(idInt).SetTitle(product.Title).SetKeywords(product.Keywords).SetDescription(product.Description).SetStatus(*product.Status).SetImage(product.Image).SetURL(product.Url).Exec(p.Context)
			if errt != nil {
				return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Product not updated, " + strings.Split(errt.Error(), ":")[3]})
			}
		}
		if err != nil {
			logs.Logger(ctx, "Update!Product not updated.", logs.ERROR)
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: 404, Message: "Product not updated"})
		}
		product.Image = selectId.Image
		return ctx.Status(fiber.StatusOK).JSON(
			response.SuccessResponse{StatusCode: 200, Message: "Image Updated.", Data: product},
		)
	}
	logs.Logger(ctx, "Store!Bad request , validate error.", logs.ERROR)

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse{StatusCode: 422, Message: validateError})
}

func (p ControllerProduct) CommentOwnProducts(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, convertError := strconv.Atoi(id)

	if convertError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{StatusCode: 400, Message: "Bad Request , Invalid type error. Type must int"})
	}
	product, userError := p.Client.Product.Query().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("id", idInt))
	}).First(p.Context)
	if userError != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{StatusCode: fiber.StatusNotFound, Message: "Comment not finding"})
	}
	setError := p.Client.Comment.Update().SetOwnerID(idInt).SetOwner(product).Exec(p.Context)
	if setError != nil {
		return ctx.Status(fiber.StatusNoContent).JSON(response.ErrorResponse{StatusCode: fiber.StatusNoContent, Message: "Category own products not creating"})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{StatusCode: fiber.StatusOK, Message: "Category products create"})
}
