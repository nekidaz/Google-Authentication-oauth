package controllers

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/nekidaz/Google-Authentication-oauth/initializers"
	"github.com/nekidaz/Google-Authentication-oauth/internal/models"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"strings"
)

var state string

func Login(c *fiber.Ctx) error {
	url := initializers.Oauth.AuthCodeURL(state)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func Callback(c *fiber.Ctx) error {
	code := c.Query("code")
	tok, err := initializers.Oauth.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	client := initializers.Oauth.Client(oauth2.NoContext, tok)
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

	var user models.User

	err = initializers.DB.Where("email = ?", userInfo.Email).First(&user).Error
	if err != nil {
		user = models.User{
			Name:  userInfo.Name,
			Email: userInfo.Email,
		}
		err = initializers.DB.Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
	} else {
		user.Name = userInfo.Name
		err = initializers.DB.Save(&user).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + tok.AccessToken,
		HTTPOnly: true,
	})

	return c.Redirect("/user", http.StatusTemporaryRedirect)
}
func ValidateGoogleToken(tokenString string) bool {
	token := oauth2.Token{
		AccessToken: tokenString,
	}

	_, err := initializers.Oauth.TokenSource(context.Background(), &token).Token()
	if err != nil {
		return false
	}

	if token.Valid() {
		return true
	}

	return false
}

func ExtractTokenFromHeader(token string) string {
	token = strings.Replace(token, "Bearer", "", 1)
	return strings.TrimSpace(token)
}
