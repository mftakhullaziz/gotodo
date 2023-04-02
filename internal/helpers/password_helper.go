package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	// Generate the bcrypt hash from the password string
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashedPassword := string(hashedPasswordBytes)
	return hashedPassword
}
