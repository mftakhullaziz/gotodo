package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"gotodo/internal/helpers"
	"net/http"
	"time"
)

func GenerateJWTToken() (string, time.Time, error) {
	// Set the expiration time of the token
	expirationTime := time.Now().Add(2 * time.Minute) // 1 day

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

func AuthenticateWithInTokenToken(tokenString string) (bool, error) {
	log := helpers.LoggerParent()
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

func MakeAuthenticatedRequest(token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
