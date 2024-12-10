package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (ar *ProductRecommendationRoutesImpl) ProductRecommendationRoutes(app *fiber.App) {
	app.Post("/recommendations",  ar.productRecommendationHandler.Create)
	app.Get("/recommendations", ar.productRecommendationHandler.GetAll)
	app.Get("/recommendations/:id", ar.productRecommendationHandler.GetById)
	app.Patch("/recommendations/:id",  ar.productRecommendationHandler.Update)
	app.Delete("/recommendations/:id",  ar.productRecommendationHandler.Delete)

}
