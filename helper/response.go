package helper

import (
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
)

func Response(ctx fiber.Ctx, config *viper.Viper, code string, message any, data any) *WebResponse[interface{}] {
	res := new(WebResponse[interface{}])
	res.Code = code
	res.Message = message
	res.Data = data

	if ctx.OriginalURL() == "" {
		// err := Loki(ctx, config, message)
		// if err != nil {
		// 	res.Code = 500
		// 	res.Message = err.Error()
		// 	res.Data = EmptyObject()
		// }
	}

	return res
}

func ResponseLimitOffset(ctx fiber.Ctx, config *viper.Viper, code string, message any, total_page any, data any) *ResponseLimit[interface{}] {
	res := new(ResponseLimit[interface{}])
	res.Code = code
	res.Message = message
	res.TotalPage = total_page
	res.Data = data

	if ctx.OriginalURL() == "" {
		// err := Loki(ctx, config, message)
		// if err != nil {
		// 	res.Code = 500
		// 	res.Message = err.Error()
		// 	res.Data = EmptyObject()
		// }
	}

	return res
}
