package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func InterfaceToMap(i interface{}) (map[string]interface{}, bool) {
	// Check if the underlying type of i is a map with string keys and interface{} values
	m, ok := i.(map[string]interface{})
	return m, ok
}

func ValueMapToInt(val interface{}) (int, error) {
	if v, ok := val.(int); ok {
		return v, nil
	}
	return 0, fmt.Errorf("value is not of type int")
}

func ValueToInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case float64:
		return int(v), nil
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case string:
		// Assuming the string represents an integer value
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return i, nil
	default:
		return 0, fmt.Errorf("conversion to int not supported for type %v", reflect.TypeOf(value))
	}
}
