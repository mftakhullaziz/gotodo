package persistence

import (
	"github.com/stretchr/testify/assert"
	"gotodo/internal/persistence/record"
	"reflect"
	"testing"
	"time"
)

func TestAccountRecord(t *testing.T) {
	t.Run("Test Account Record TableName", func(t *testing.T) {
		var account record.AccountRecord
		assert.NotNil(t, account)
		assert.Equal(t, "accounts", account.TableName())
	})

	t.Run("Test Account Record Struct Field Values", func(t *testing.T) {
		now := time.Now().UTC()
		account := record.AccountRecord{
			UserID:    1,
			Username:  "user",
			Password:  "password",
			Status:    "active",
			CreatedAt: now,
			UpdatedAt: now,
		}
		assert.Equal(t, "user", account.Username)
		assert.Equal(t, "password", account.Password)
		assert.Equal(t, "active", account.Status)
		assert.Equal(t, now, account.CreatedAt)
		assert.Equal(t, now, account.UpdatedAt)
	})

	t.Run("Test Account Record Struct Field DataType", func(t *testing.T) {
		var account record.AccountRecord

		accountIdType := reflect.TypeOf(account.AccountID).Kind().String()
		assert.Equal(t, "uint", accountIdType)

		userIdType := reflect.TypeOf(account.UserID).Kind().String()
		assert.Equal(t, "int", userIdType)

		usernameType := reflect.TypeOf(account.Username).Kind().String()
		assert.Equal(t, "string", usernameType)

		passwordType := reflect.TypeOf(account.Password).Kind().String()
		assert.Equal(t, "string", passwordType)

		statusType := reflect.TypeOf(account.Status).Kind().String()
		assert.Equal(t, "string", statusType)

		createAtType := reflect.TypeOf(account.CreatedAt).String()
		assert.Equal(t, "time.Time", createAtType)

		updateAtType := reflect.TypeOf(account.UpdatedAt).String()
		assert.Equal(t, "time.Time", updateAtType)
	})
}
