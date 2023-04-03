package helpers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

func Value(a interface{}) (driver.Value, error) {
	return json.Marshal(a)
}

func Scan(value interface{}, a *interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type")
	}
	return json.Unmarshal(bytes, a)
}
