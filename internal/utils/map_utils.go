package utils

func InterfaceToMap(i interface{}) (map[string]interface{}, bool) {
	// Check if the underlying type of i is a map with string keys and interface{} values
	m, ok := i.(map[string]interface{})
	return m, ok
}
