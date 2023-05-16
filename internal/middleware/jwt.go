package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"gotodo/internal/utils"
	"net/http"
	"time"
)

func GenerateJWTToken() (string, time.Time, error) {
	// Set the expiration time of the token
	expirationTime := time.Now().Add(1 * time.Hour) // 1 day
	utils.LoggerParent().Infoln("expire token time at: ", expirationTime)
	// Create a new JWT token with the claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := "gotodo-secret-key"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

// AuthenticateWithInToken function to check token if available or not using JWT and check token if expire or not
func AuthenticateWithInToken(tokenString string) (bool, error) {
	log := utils.LoggerParent()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Here you should return the key used to sign the token
		// For example, you can return a []byte containing a secret key
		return []byte("gotodo-secret-key"), nil
	})
	if err != nil {
		return false, err
	}

	// Check if the token has expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := claims["exp"].(float64)
		log.Infoln("expiresAt: ", exp)
		if exp < float64(time.Now().Unix()) {
			return false, nil
		}
		return true, nil
	}

	return false, nil
}

func MakeAuthenticatedRequest(token string) error {
	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	return nil
	/*
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		return nil
	*/
}
