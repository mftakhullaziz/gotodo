package database

import (
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB, models ...interface{}) error {
	migrator := db.Migrator()

	for _, model := range models {
		if migrator.HasTable(model) {
			continue
		}
		err := db.AutoMigrate(model)
		if err != nil {
			panic(err.Error())
		}
	}
	return nil
}
