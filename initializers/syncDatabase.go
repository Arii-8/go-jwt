package initializers

import "go-jwt/models"

// sync database auto migrate
func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
