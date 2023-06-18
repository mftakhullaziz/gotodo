package middleware

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

var c *cache.Cache

// Initialize the cache when the application starts up
func init() {
	c = cache.New(1*time.Hour, 24*time.Hour)
}

// GenerateTokenToCache Generate a new token for the user and store it in the cache
func GenerateTokenToCache(userID string, token string, expirationTime time.Time) error {
	//log := utils.LoggerParent()
	//log.Log.Infoln("cache token expire_at countdown: ", expirationTime.Sub(time.Now()))
	err := c.Add(token, userID, expirationTime.Sub(time.Now()))
	if err != nil {
		return err
	}
	return nil
}

// AuthenticateUser Authenticate the user by checking their token against the cache and check token if expire or not
func AuthenticateUser(token string) (string, error) {
	userID, _, err := c.GetWithExpiration(token)

	if err != true {
		return "", errors.New("user account unauthorized")
	}

	// Update the expiration time of the token to extend the user's session
	c.Set(token, userID, 30*time.Minute)
	return userID.(string), nil
}

func CheckAndRemoveTokenFromCache(token string) error {
	// Check if the token exists in the cache
	_, exists := c.Get(token)
	if exists {
		// Remove the token from the cache
		c.Delete(token)
	}

	return nil
}
