package mymongo

import (
    "context"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
		"github.com/sirupsen/logrus"
)

var (
	ctx = context.TODO()
	clientOptions *options.ClientOptions
	client *mongo.Client
	err error
)

func getClient() (*mongo.Client, error) {

	if client != nil {
		return client, nil
	}
	
	// Set client options
	clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/?connect=direct")

	// Connect to MongoDB
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
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
