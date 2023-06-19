package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekidaz/Google-Authentication-oauth/internal/models"
)

func Profile(c *fiber.Ctx) error {
	user := c.Context().UserValue("user")
	return c.Render("profile", user.(*models.User))
}
