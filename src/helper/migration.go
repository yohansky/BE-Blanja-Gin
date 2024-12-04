package helper

import (
	"be-blanja/src/config"
	"be-blanja/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
	)
}
