package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	errs "gotodo/internal/utils/errors"
	"gotodo/internal/utils/logger"
	"os"
	"time"
)

func NewDatabaseConnection(ctx context.Context, path string) (db *gorm.DB, error error) {
	log := logger.LoggerParent()

	err := godotenv.Load(path)
	errs.FatalIfErrorWithCustomMessage(err, log, "Error loading .env.test file: %v")

	// Do get from environment file
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	hostname := os.Getenv("MYSQL_HOST")
	databaseName := os.Getenv("MYSQL_NAME")

	databaseConnection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, hostname, databaseName)

	connection, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{})
	errs.PanicIfErrorWithCustomMessage(err, "Failed to create connection to database")

	sqlDB, err := connection.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.PingContext(ctx)

	if err != nil {
		_ = fmt.Errorf("failed to ping database: %w", err)
	}

	return connection, nil
}
