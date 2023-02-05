package initializers

import (
	"API_go/go_test/models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
