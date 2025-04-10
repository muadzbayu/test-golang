package article

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type ArticleHandler interface {
	CreateData(ctx fiber.Ctx) error
	GetDataLimitOffset(ctx fiber.Ctx) error
	GetDataById(ctx fiber.Ctx) error
	EditData(ctx fiber.Ctx) error
	DeleteData(ctx fiber.Ctx) error
	TrashData(ctx fiber.Ctx) error
}

type articleHandler struct {
	Log     *logrus.Logger
	UseCase ArticleUseCase
}

func NewArticleHandler(useCase ArticleUseCase) ArticleHandler {
	return &articleHandler{
		UseCase: useCase,
	}
}

func (c *articleHandler) CreateData(ctx fiber.Ctx) error {
	res := c.UseCase.CreateData(ctx)
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *articleHandler) GetDataLimitOffset(ctx fiber.Ctx) error {
	res := c.UseCase.GetDataLimitOffset(ctx)
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *articleHandler) GetDataById(ctx fiber.Ctx) error {
	res := c.UseCase.GetDataById(ctx)
	return ctx.Status(fiber.StatusOK).JSON(res)
}
func (c *articleHandler) EditData(ctx fiber.Ctx) error {
	res := c.UseCase.EditData(ctx)
	return ctx.Status(fiber.StatusOK).JSON(res)
}
func (c *articleHandler) DeleteData(ctx fiber.Ctx) error {
	res := c.UseCase.DeleteData(ctx)
	return ctx.Status(fiber.StatusOK).JSON(res)
}
func (c *articleHandler) TrashData(ctx fiber.Ctx) error {
	res := c.UseCase.TrashData(ctx)
	return ctx.Status(fiber.StatusOK).JSON(res)
}
