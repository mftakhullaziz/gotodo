package persistence

import (
	"github.com/stretchr/testify/assert"
	"gotodo/internal/infra/persistence/record"
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
}
