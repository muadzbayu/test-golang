package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/muadzbayu/test-golang/app/article"
	"github.com/muadzbayu/test-golang/route"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	DB       *gorm.DB
	App      *fiber.App
	Config   *viper.Viper
	Validate *validator.Validate
}

func AllConfig(config *Config) {
	config.App.Get("/", func(c fiber.Ctx) error {
		return c.JSON("welcome!")
	})

	// SETTLEMENT
	articleRepository := article.NewArticleRepository()
	articleUseCase := article.NewArticleUseCase(config.DB, config.Validate, articleRepository, config.Config)
	articleHandler := article.NewArticleHandler(articleUseCase)

	routeConfig := route.RouteConfig{
		App:            config.App,
		ArticleHandler: articleHandler,
	}

	routeConfig.Setup()

	config.App.Get("*", func(c fiber.Ctx) error {
		return c.JSON("what are you looking for")
	})
}
