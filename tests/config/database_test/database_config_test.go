package database_test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gotodo/config/database"
	"os"
	"path/filepath"
	"testing"
)

func TestNewDatabaseConnection(t *testing.T) {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	rootDir := filepath.Dir(filepath.Dir(filepath.Dir(dir)))
	assert.NotNil(t, rootDir)

	// Construct the full path to the .env file
	envPath := filepath.Join(rootDir, "gotodo", ".env")
	assert.NotNil(t, envPath)

	err = godotenv.Load(envPath)
	assert.NoError(t, err)

	// Call the function being tested
	db, err := database.NewDatabaseConnection(context.Background())

	// Make assertions on the results
	assert.NoError(t, err, "Unexpected error creating database connection")
	assert.NotNil(t, db, "Expected non-nil database connection")
	assert.Equal(t, "*gorm.DB", fmt.Sprintf("%T", db), "Unexpected type for database connection")
	assert.Equal(t, "mysql", db.Dialect().GetName(), "Unexpected dialect for database connection")

	// Close the database connection
	err = db.Close()
	assert.NoError(t, err, "Unexpected error closing database connection")
}
