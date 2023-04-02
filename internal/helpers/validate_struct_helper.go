package helpers

import "github.com/badoux/checkmail"

func ValidateEmail(email string) error {
	if err := checkmail.ValidateFormat(email); err != nil {
		return err
	}
	return nil
}
