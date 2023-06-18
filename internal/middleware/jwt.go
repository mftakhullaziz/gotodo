package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gotodo/internal/utils"
	"net/http"
	"time"
)

func GenerateJWTToken() (string, time.Time, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix() // Issued At (current time in UNIX timestamp format)

	// Set the expiration time for the token (optional)
	expirationTime := time.Now().Add(15 * time.Second)
	claims["exp"] = expirationTime.Unix()

	// Set the JWT secret key
	// Replace "your-secret-key" with your own secret key
	secretKey := []byte("gotodo-secret-key")

	// Generate the JWT token string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to generate JWT token: %v", err)
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
		log.Log.Infoln("expiresAt: ", exp)
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

func CheckAndRemoveExpiredToken(tokenString string) (bool, error) {
	// Parse the JWT token string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Set the JWT secret key
		// Replace "your-secret-key" with your own secret key
		secretKey := []byte("gotodo-secret-key")
		return secretKey, nil
	})

	if err != nil {
		return false, fmt.Errorf("failed to parse JWT token: %v", err)
	}

	// Check if the token is valid and has not expired
	if token.Valid {
		// Token is valid
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return false, fmt.Errorf("invalid JWT claims")
		}

		// Check if the "exp" claim exists and if the token has expired
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			currentTime := time.Now()
			if currentTime.After(expirationTime) {
				// Token has expired
				return false, nil
			}
		}
		return false, nil
	} else {
		return false, nil
	}

	// Token is invalid or has some other error
	return true, fmt.Errorf("invalid JWT token")
}
