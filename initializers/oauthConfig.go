package initializers

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

var Oauth *oauth2.Config

func InitializeOAuthConfig() {
	clientID := os.Getenv("ClientID")
	clientSecret := os.Getenv("ClientSecret")
	redirectURL := os.Getenv("RedirectURL")

	if clientID == "" || clientSecret == "" || redirectURL == "" {
		log.Fatal("Invalid .env variables")
	}

	Oauth = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

}
