package persistence

import (
	"github.com/stretchr/testify/assert"
	"gotodo/internal/persistence/record"
	"testing"
	"time"
)

func TestUserDetail(t *testing.T) {
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

	// Test TableName() function
	assert.Equal(t, "user_details", userDetail.TableName())

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
}
