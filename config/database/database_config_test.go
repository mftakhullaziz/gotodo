package database

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gotodo/config"
	"testing"
)

func TestNewDatabaseConnection(t *testing.T) {
	nameEnv := config.LoadEnv(".env.test")
	assert.NotNil(t, nameEnv)

	// Call the function being tested
	db, err := NewDatabaseConnection(context.Background(), nameEnv)

	// Make assertions on the results
	assert.NoError(t, err, "Unexpected errors creating database connection")
	assert.NotNil(t, db, "Expected non-nil database connection")
	assert.Equal(t, "*gorm.DB", fmt.Sprintf("%T", db), "Unexpected type for database connection")
	assert.Equal(t, "mysql", db.Name(), "Unexpected dialect for database connection")
}
