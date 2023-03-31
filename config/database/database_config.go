package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"time"
)

// NewDatabaseConnection
// Do: Function to open connection with database mysql
// Param: Context
func NewDatabaseConnection(ctx context.Context) (db *gorm.DB, errs error) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rootDir := filepath.Dir(filepath.Dir(filepath.Dir(dir)))

	// Construct the full path to the .env file
	envPath := filepath.Join(rootDir, ".env")
	err = godotenv.Load(envPath)

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Do get from environment file
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	hostname := os.Getenv("MYSQL_HOST")
	databaseName := os.Getenv("MYSQL_NAME")

	databaseConnection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, hostname, databaseName)

	//connection, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{})

	connection, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{})

	if err != nil {
		panic("Failed to create connection to database")
	}

	sqlDB, err := connection.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.PingContext(ctx)

	if err != nil {
		_ = fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to database")

	return connection, nil
}
