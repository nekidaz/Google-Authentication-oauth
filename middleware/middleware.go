package middlewares

import (
	"GoogleAuthv2.0/controllers"
	"GoogleAuthv2.0/initializers"
	"GoogleAuthv2.0/internal/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func RequireAuth(c *fiber.Ctx) error {
	authorizationHeader := c.Cookies("Authorization")

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

	token := controllers.ExtractTokenFromHeader(authorizationHeader)

	client := initializers.Oauth.Client(oauth2.NoContext, &oauth2.Token{AccessToken: token})
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		Name          string `json:"name"`
	}

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		log.Fatal(err)
	}
	user := &models.User{
		Name:  userInfo.Name,
		Email: userInfo.Email,
	}
	c.Context().SetUserValue("user", user)
	return c.Next()
}
