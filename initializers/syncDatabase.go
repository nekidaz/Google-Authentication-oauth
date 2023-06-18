package initializers

import "GoogleAuthv2.0/internal/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
