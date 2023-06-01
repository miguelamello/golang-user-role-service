package myredis

import (
	"encoding/json"

	"github.com/miguelamello/user-domain-role-service/graph/model"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	//"github.com/miguelamello/user-domain-role-service/graph/packages/mymongo"
)

// GetUser retrieve the user object to Redis
func GetUser(userID string) (*model.User, error) {

	// Get the Redis client
	client, err := getClient()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// Retrieve the user data from Redis
	userJSON, err := client.Do(ctx, "JSON.GET", userID).Text()
	if err != nil {
		if err == redis.Nil {
			logrus.Error(err)
			return nil, err
		}
		logrus.Error(err)
		return nil, err
	}

	// Parse the user JSON into a User struct
	var user model.User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &user, nil

}

// SaveUser saves the user object to Redis
func SaveUser(user *model.User) (bool, error) {

	// Convert the user object to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	// Get the Redis client
	client, err := getClient()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	// Save the user object in Redis
	err = client.Do(ctx, "JSON.SET", user.ID, "$", string(userJSON)).Err()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	//go mymongo.PushUser(user)
	return true, nil

}
