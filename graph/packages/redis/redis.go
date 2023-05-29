package redis

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/miguelamello/user-domain-role-service/graph/model"
)

var (
	ctx    = context.Background()
	client *redis.Client
)

// Initializes the Redis client
func getClient() (*redis.Client, error) {

	if client != nil {
		return client, nil
	}

	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil

}

// GetUser retrieve the user object to Redis
func GetUser(userID string) (*model.User, error) {

	// Get the Redis client
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	// Retrieve the user data from Redis
	userJSON, err := client.Get(ctx, userID).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, err
		}
		return nil, err
	}

	// Parse the user JSON into a User struct
	var user model.User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

// SaveUser saves the user object to Redis
func SaveUser(user *model.User) (bool, error) {

	// Convert the user object to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return false, err
	}

	// Get the Redis client
	client, err := getClient()
	if err != nil {
		return false, err
	}

	// Save the user object in Redis
	err = client.Set(ctx, user.ID, string(userJSON), 0).Err()
	if err != nil {
		return false, err
	}

	return true, nil

}
