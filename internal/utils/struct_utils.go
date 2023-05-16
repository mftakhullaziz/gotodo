package utils

import (
	"gotodo/internal/persistence/record"
	"reflect"
)

func HasValue(s interface{}) bool {
	// Get the value of the struct
	v := reflect.ValueOf(s)
	// Loop over the struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// Check if the field's value is non-zero
		if !field.IsZero() {
			return true
		}
	}
	return false
}

type UserAccounts struct {
	Accounts record.AccountRecord `gorm:"foreignKey:UserID;references:UserID"`
	Users    record.UserDetailRecord
}

func HasValueSlice(handler interface{}) bool {
	v := reflect.ValueOf(handler)
	return v.Kind() != reflect.Ptr && !v.IsNil()
}
