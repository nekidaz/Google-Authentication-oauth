package controllers

import (
	"GoogleAuthv2.0/internal/models"
	"github.com/gofiber/fiber/v2"
)

func Profile(c *fiber.Ctx) error {
	user := c.Context().UserValue("user")
	return c.Render("profile", user.(*models.User))
}
