package initializers

import "github.com/nekidaz/Google-Authentication-oauth/internal/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
