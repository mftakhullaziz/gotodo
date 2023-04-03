package middleware

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"gotodo/internal/helpers"
	"time"
)

var c *cache.Cache

// Initialize the cache when the application starts up
func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

// GenerateTokenToCache Generate a new token for the user and store it in the cache
func GenerateTokenToCache(userID string, token string, expirationTime time.Time) error {
	log := helpers.LoggerParent()
	log.Infoln("Cache token expire at: ", expirationTime.Sub(time.Now()))

	err := c.Add(token, userID, expirationTime.Sub(time.Now()))
	if err != nil {
		return err
	}
	return nil
}

// AuthenticateUser Authenticate the user by checking their token against the cache
func AuthenticateUser(token string) (string, error) {
	userID, found := c.Get(token)
	if !found {
		return "", errors.New("user account unauthorized")
	}

	// Update the expiration time of the token to extend the user's session
	c.Set(token, userID, 30*time.Minute)
	return userID.(string), nil
}
