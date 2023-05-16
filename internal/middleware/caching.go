package middleware

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"gotodo/internal/utils"
	"time"
)

var c *cache.Cache

// Initialize the cache when the application starts up
func init() {
	c = cache.New(1*time.Hour, 24*time.Hour)
}

// GenerateTokenToCache Generate a new token for the user and store it in the cache
func GenerateTokenToCache(userID string, token string, expirationTime time.Time) error {
	log := utils.LoggerParent()
	log.Infoln("cache token expire_at countdown: ", expirationTime.Sub(time.Now()))

	err := c.Add(token, userID, expirationTime.Sub(time.Now()))
	if err != nil {
		return err
	}
	return nil
}

// AuthenticateUser Authenticate the user by checking their token against the cache and check token if expire or not
func AuthenticateUser(token string) (string, error) {
	userID, found, err := c.GetWithExpiration(token)
	fmt.Println("Tokens: ", found)
	if err != true {
		return "", errors.New("user account unauthorized")
	}

	// Update the expiration time of the token to extend the user's session
	c.Set(token, userID, 30*time.Minute)
	return userID.(string), nil
}
