package database

import (
	"gorm.io/gorm"
	"gotodo/internal/utils"
)

func MigrateDatabase(db *gorm.DB, models ...interface{}) error {
	migrator := db.Migrator()

	for _, model := range models {
		if migrator.HasTable(model) {
			continue
		}

		err := db.AutoMigrate(model)
		utils.LoggerIfError(err)
	}
	return nil
}
