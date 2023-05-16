package database

import (
	"gorm.io/gorm"
	errs "gotodo/internal/utils/errors"
)

func MigrateDatabase(db *gorm.DB, models ...interface{}) error {
	migrator := db.Migrator()

	for _, model := range models {
		if migrator.HasTable(model) {
			continue
		}

		err := db.AutoMigrate(model)
		errs.LoggerIfError(err)
	}
	return nil
}
