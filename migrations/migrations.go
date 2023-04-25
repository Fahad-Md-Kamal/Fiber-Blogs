package migrations

import (
	"github.com/fahad-md-kamal/fiber-blogs/database"
	usermodels "github.com/fahad-md-kamal/fiber-blogs/users/models"
)

func MigrateChanges() {
	database.DB.AutoMigrate(
		&usermodels.Users{},
	)
}
