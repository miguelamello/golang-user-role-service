package myredis

import (
	"encoding/json"

	"github.com/miguelamello/user-domain-role-service/graph/model"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/miguelamello/user-domain-role-service/graph/packages/mymongo"
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
	var user model.User
	userJSON, err := client.Do(ctx, "JSON.GET", userID).Text()
	if err != nil {
		if err == redis.Nil {
			//logrus.Error(1, err)
			item, err := mymongo.PullUser(userID)
			if err != nil {
				logrus.Error(err)
				return nil, err
			}
			user = model.User{
				ID: item[3].Value.(string),
				Name: item[1].Value.(string),
				Email: item[2].Value.(string),
			}
			go SyncUser(&user)
			return &user, nil
		}
		logrus.Error(2, err)
		return nil, err
	}

	// Parse the user JSON into a User struct
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		logrus.Error(3, err)
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

	go mymongo.PushUser(userJSON)
	return true, nil

}

// SyncUser sincronize the user object to Redis
func SyncUser(user *model.User) {

	// Convert the user object to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		logrus.Error(err)
	}

	// Get the Redis client
	client, err := getClient()
	if err != nil {
		logrus.Error(err)
	}

	// Save the user object in Redis
	err = client.Do(ctx, "JSON.SET", user.ID, "$", string(userJSON)).Err()
	if err != nil {
		logrus.Error(err)
	}

}

