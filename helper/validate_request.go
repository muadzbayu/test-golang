package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func ValidateRequest(ctx fiber.Ctx, c *validator.Validate, request any) error {

	err := ctx.Bind().Body(request)
	if err != nil {
		return err
	}

	// Validate struct
	return c.Struct(request)
}
