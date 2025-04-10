package article

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/muadzbayu/test-golang/helper"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ArticleUseCase interface {
	CreateData(ctx fiber.Ctx) *helper.WebResponse[interface{}]
	GetDataLimitOffset(ctx fiber.Ctx) *helper.ResponseLimit[interface{}]
	GetDataById(ctx fiber.Ctx) *helper.WebResponse[interface{}]
	EditData(ctx fiber.Ctx) *helper.WebResponse[interface{}]
	DeleteData(ctx fiber.Ctx) *helper.WebResponse[interface{}]
	TrashData(ctx fiber.Ctx) *helper.WebResponse[interface{}]
}

type articleUseCase struct {
	DB                *gorm.DB
	Validate          *validator.Validate
	ArticleRepository ArticleRepository
	Config            *viper.Viper
}

func NewArticleUseCase(db *gorm.DB, validate *validator.Validate, ArticleRepository ArticleRepository, viper *viper.Viper) ArticleUseCase {
	return &articleUseCase{
		DB:                db,
		Validate:          validate,
		ArticleRepository: ArticleRepository,
		Config:            viper,
	}
}

func (c *articleUseCase) CreateData(ctx fiber.Ctx) *helper.WebResponse[interface{}] {
	fmt.Println("~CREATE DATA~")
	response := &helper.WebResponse[interface{}]{}
	request := new(ArticleRequestV2)

	tx := c.DB.WithContext(ctx.Context())
	c.Validate = validator.New()

	err := helper.ValidateRequest(ctx, c.Validate, request)
	if err != nil {
		response = helper.Response(ctx, c.Config, "02", err.Error(), nil)

		return response
	}

	// Check ketersediaan username di DB
	result, err := c.ArticleRepository.CreateData(tx, ArticleData(*request))
	if err != nil {
		return helper.Response(ctx, c.Config, "02", err.Error(), nil)

	}

	if result == 0 {
		return helper.Response(ctx, c.Config, "02", "Add new data has failed", nil)
	}

	response = helper.Response(ctx, c.Config, "00", "Sukses", nil)

	return response
}

func (c *articleUseCase) GetDataLimitOffset(ctx fiber.Ctx) *helper.ResponseLimit[interface{}] {
	fmt.Println("~GET DATA LIMIT OFFSET~")
	request := new(ArticleLimitOffset)

	// Convert path params to int
	limitStr := ctx.Params("limit")
	offsetStr := ctx.Params("offset")

	// check preview
	isPreview := ctx.Query("preview") == "true"

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return helper.ResponseLimitOffset(ctx, c.Config, "02", "Limit must be a number", 0, nil)
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return helper.ResponseLimitOffset(ctx, c.Config, "02", "Offset must be a number", 0, nil)
	}

	// Assign to request struct
	request.Limit = limit
	request.Offset = offset

	// Validate using validator
	c.Validate = validator.New()
	err = c.Validate.Struct(request)
	if err != nil {
		return helper.ResponseLimitOffset(ctx, c.Config, "02", err.Error(), 0, nil)
	}

	tx := c.DB.WithContext(ctx.Context())

	result, total_page, err := c.ArticleRepository.GetDataLimitOffset(tx, request.Limit, request.Offset, isPreview)
	if err != nil {
		return helper.ResponseLimitOffset(ctx, c.Config, "03", err.Error(), 0, nil)
	}

	return helper.ResponseLimitOffset(ctx, c.Config, "00", "Sukses", total_page, result)

}

func (c *articleUseCase) GetDataById(ctx fiber.Ctx) *helper.WebResponse[interface{}] {
	fmt.Println("~GET DATA BY ID~")
	response := &helper.WebResponse[interface{}]{}
	request := new(ArticleId)

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", "Offset must be a number", nil)
	}

	request.Id = id

	tx := c.DB.WithContext(ctx.Context())
	c.Validate = validator.New()

	err = c.Validate.Struct(request)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", err.Error(), nil)
	}

	// Check ketersediaan username di DB
	result, err := c.ArticleRepository.GetDataById(tx, request.Id)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", err.Error(), nil)

	}

	response = helper.Response(ctx, c.Config, "00", "Sukses", result)

	return response
}

func (c *articleUseCase) EditData(ctx fiber.Ctx) *helper.WebResponse[interface{}] {
	fmt.Println("~EDIT ARTICLE~")
	request := new(ArticleRequestV2)

	// Get the ID from params and convert to int
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", "Invalid ID parameter", nil)
	}

	// Parse and validate the body
	tx := c.DB.WithContext(ctx.Context())
	c.Validate = validator.New()

	err = helper.ValidateRequest(ctx, c.Validate, request)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", err.Error(), nil)
	}

	// Convert request into the type expected by repository
	data := ArticleData{
		Title:    request.Title,
		Content:  request.Content,
		Category: request.Category,
		Status:   request.Status,
	}

	result, err := c.ArticleRepository.EditData(tx, id, data)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", err.Error(), nil)
	}

	if result == 0 {
		return helper.Response(ctx, c.Config, "02", "Edit data has failed", nil)
	}

	return helper.Response(ctx, c.Config, "00", "Success", nil)
}

func (c *articleUseCase) DeleteData(ctx fiber.Ctx) *helper.WebResponse[interface{}] {
	// Convert string ID from path to int
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", "Invalid ID parameter", nil)
	}

	tx := c.DB.WithContext(ctx.Context())

	result, err := c.ArticleRepository.DeleteData(tx, id)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", err.Error(), nil)
	}

	if result == 0 {
		return helper.Response(ctx, c.Config, "02", "Delete data has failed", nil)
	}

	return helper.Response(ctx, c.Config, "00", "Sukses", result)
}

func (c *articleUseCase) TrashData(ctx fiber.Ctx) *helper.WebResponse[interface{}] {
	fmt.Println("~TRASH ARTICLE~")

	// Get the ID from params and convert to int
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", "Invalid ID parameter", nil)
	}

	// Parse and validate the body
	tx := c.DB.WithContext(ctx.Context())

	result, err := c.ArticleRepository.TrashData(tx, id)
	if err != nil {
		return helper.Response(ctx, c.Config, "02", err.Error(), nil)
	}

	if result == 0 {
		return helper.Response(ctx, c.Config, "02", "Delete data has failed", nil)
	}

	return helper.Response(ctx, c.Config, "00", "Success", nil)
}
