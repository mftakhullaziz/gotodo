package persistence

import (
	"github.com/stretchr/testify/assert"
	"gotodo/internal/persistence/record"
	"reflect"
	"testing"
	"time"
)

func TestUserDetail(t *testing.T) {
	t.Run("Test User Detail Record TableName", func(t *testing.T) {
		// Test TableName() function
		var userDetail record.UserDetailRecord
		assert.NotNil(t, userDetail)
		assert.Equal(t, "user_details", userDetail.TableName(), "Matcher")
	})

	t.Run("Test User Detail Struct Field Value", func(t *testing.T) {
		now := time.Now().UTC()
		assert.NotNil(t, now)

		userDetail := record.UserDetailRecord{
			Username:    "username",
			Password:    "password",
			Email:       "test@email.com",
			Name:        "name",
			MobilePhone: 12345678999999,
			Address:     "jakarta",
			Status:      "active",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// Test struct field values
		assert.Equal(t, "username", userDetail.Username)
		assert.Equal(t, "password", userDetail.Password)
		assert.Equal(t, "test@email.com", userDetail.Email)
		assert.Equal(t, "name", userDetail.Name)
		assert.Equal(t, 12345678999999, userDetail.MobilePhone)
		assert.Equal(t, "jakarta", userDetail.Address)
		assert.Equal(t, "active", userDetail.Status)
		assert.Equal(t, now, userDetail.CreatedAt)
		assert.Equal(t, now, userDetail.UpdatedAt)
	})

	t.Run("Test User Detail DataType", func(t *testing.T) {
		var userDetail record.UserDetailRecord

		idType := reflect.TypeOf(userDetail.UserID).Kind().String()
		assert.Equal(t, "uint", idType)

		usernameType := reflect.TypeOf(userDetail.Username).Kind().String()
		assert.Equal(t, "string", usernameType)

		passwordType := reflect.TypeOf(userDetail.Password).Kind().String()
		assert.Equal(t, "string", passwordType)

		emailType := reflect.TypeOf(userDetail.Email).Kind().String()
		assert.Equal(t, "string", emailType)

		nameType := reflect.TypeOf(userDetail.Name).Kind().String()
		assert.Equal(t, "string", nameType)

		phoneType := reflect.TypeOf(userDetail.MobilePhone).Kind().String()
		assert.Equal(t, "int", phoneType)

		addressType := reflect.TypeOf(userDetail.Address).Kind().String()
		assert.Equal(t, "string", addressType)

		statusType := reflect.TypeOf(userDetail.Status).Kind().String()
		assert.Equal(t, "string", statusType)

		createAtType := reflect.TypeOf(userDetail.CreatedAt).String()
		assert.Equal(t, "time.Time", createAtType)

		updateAtType := reflect.TypeOf(userDetail.UpdatedAt).String()
		assert.Equal(t, "time.Time", updateAtType)
	})
}
