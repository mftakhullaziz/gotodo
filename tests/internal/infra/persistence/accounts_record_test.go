package persistence

import (
	"github.com/stretchr/testify/assert"
	"gotodo/internal/persistence/record"
	"testing"
	"time"
)

func TestAccountRecord(t *testing.T) {
	now := time.Now().UTC()

	account := record.AccountRecord{
		UserID:    1,
		Username:  "root_name",
		Password:  "password",
		Status:    "active",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Test TableName() function
	assert.Equal(t, "accounts", account.TableName())

	// Test struct field values
	assert.Equal(t, "root_name", account.Username)
	assert.Equal(t, "password", account.Password)
	assert.Equal(t, "active", account.Status)
	assert.Equal(t, now, account.CreatedAt)
	assert.Equal(t, now, account.UpdatedAt)
}
