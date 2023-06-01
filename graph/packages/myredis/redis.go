package myredis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	ctx    = context.Background()
	client *redis.Client
	options *redis.Options
)

// Initializes the Redis client
func getClient() (*redis.Client, error) {

	if client != nil {
		return client, nil
	}

	options = &redis.Options {
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}

	client = redis.NewClient(options)

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return client, nil

}

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func init() {

	initLogger() // Initializes the logger. Should send error to Log Server.
	_, err := getClient()
	if err != nil {
		logrus.Error(err)
	}

}
