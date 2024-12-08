package routes

import (
	"acne-scan-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (authRoutes *AuthRoutesImpl) AuthRoutes(app *fiber.App) {
	app.Post("/auth",authRoutes.authHandlers.Login)
	app.Post("/auth/register",middleware.SuperAdminMiddleware(),authRoutes.authHandlers.Register)
}