package middlewares

import (
	"GoogleAuthv2.0/controllers"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func RequireAuth(c *fiber.Ctx) error {

	authorizationHeader := c.Get("Authorization")
	if authorizationHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if !controllers.ValidateGoogleToken(authorizationHeader) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	return c.Next()

}
