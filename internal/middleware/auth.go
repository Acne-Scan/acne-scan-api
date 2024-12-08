package middleware

import (
	"acne-scan-api/internal/pkg/jwt"
	"acne-scan-api/internal/pkg/response"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		splittedToken := strings.Split(token, " ")

		if splittedToken[0] != "Bearer" {
			return response.BadRequest(c, "invalid token.", fmt.Errorf("invalid token"))
		}

		data, err := jwt.ExtractToken(splittedToken[1])
		if err != nil {
			return err
		}

		dataRole := strings.ToLower(data.Role)

		if dataRole != role {
			return response.StatusForbidden(c, "forbidden", err)
		}

		c.Locals("user_id", data.UserId)

		return c.Next()
	}
}

func AdminMiddleware() fiber.Handler {
	return AuthMiddleware("admin")
}

func UserMiddleware() fiber.Handler {
	return AuthMiddleware("user")
}
