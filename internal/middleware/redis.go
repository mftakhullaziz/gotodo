package middleware

import (
	"context"
	"github.com/go-redis/redis/v8"
	errs "gotodo/internal/utils/errors"
	"time"
)

func NewRedisClient(ctx context.Context) (*redis.Client, error) {
	// create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password, if any
		DB:       0,                // Redis database number
	})

	// ping the Redis server to check if it's reachable
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewSaveTokenRedis(token string, userId int64, ctx context.Context) error {
	// create a new Redis client
	client, err := NewRedisClient(ctx)
	errs.LoggerIfError(err)

	// save the token in Redis with an expiration time of 1 hour
	err = client.Set(ctx, token, userId, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func NewGetUserIDByToken(token string, ctx context.Context) (int64, error) {
	// create a new Redis client
	client, err := NewRedisClient(ctx)
	errs.LoggerIfError(err)

	// retrieve the userID from Redis
	userID, err := client.Get(ctx, token).Int64()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
