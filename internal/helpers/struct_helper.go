package helpers

import "reflect"

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
