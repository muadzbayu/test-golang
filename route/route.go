package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/muadzbayu/test-golang/app/article"
)

type RouteConfig struct {
	App            *fiber.App
	ArticleHandler article.ArticleHandler
}

func (c *RouteConfig) Setup() {
	c.ArticleRoute()
}

func (c *RouteConfig) ArticleRoute() {
	articleRoute := c.App.Group("/article")
	//create articles
	articleRoute.Post("/", c.ArticleHandler.CreateData)

	//get data article by limit and offset
	articleRoute.Get("/:limit/:offset", c.ArticleHandler.GetDataLimitOffset)

	//get data article by id
	articleRoute.Get("/:id", c.ArticleHandler.GetDataById)

	//edit article
	articleRoute.Post("/:id", c.ArticleHandler.EditData)

	//delete article
	articleRoute.Delete("/:id", c.ArticleHandler.DeleteData)

	//delete article
	articleRoute.Post("/trash/:id", c.ArticleHandler.TrashData)
}
