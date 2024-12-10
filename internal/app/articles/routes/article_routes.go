package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (articleRoutes *ArticleRoutesImpl) ArticleRoutes(app *fiber.App) {
	app.Post("/articles", articleRoutes.ArticleHandler.Create)
	app.Get("/articles", articleRoutes.ArticleHandler.GetAll)
	app.Get("/articles/:id", articleRoutes.ArticleHandler.GetById)
	app.Delete("/articles/:id", articleRoutes.ArticleHandler.Delete)
	app.Patch("/articles/:id", articleRoutes.ArticleHandler.Update)

}