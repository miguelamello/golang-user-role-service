package mymongo

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name  string
	Email string
	Redis string
}

func PullUser(redis_id string) (bson.D, error) {

	// Get the Mongo client
	client, err = getClient()
	if err != nil {
		logrus.Error(err)
	}

	// Access the target database and collection
	collection := client.Database("udrs").Collection("users")

	// Define the filter to retrieve the item
	filter := bson.M{"redis": redis_id}

	// Retrieve the User from the collection
	var result bson.D
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logrus.Error("User not found")
			return nil, err
		} else {
			logrus.Error(err)
			return nil, err
		}
	}

	return result, nil

}

func PushUser(userJSON []byte) {

	// Get the Mongo client
	client, err = getClient()
	if err != nil {
		logrus.Error(err)
	}

	// Access the target database and collection
	collection := client.Database("udrs").Collection("users")

	// Define the json object container
	var item interface{}

	// Unmarshal the JSON string into an Item struct
	err = json.Unmarshal(userJSON, &item)
	if err != nil {
		logrus.Error(err)
	}

	// Define the user object to be inserted
	user := User{
		Name:  item.(map[string]interface{})["name"].(string),
		Email: item.(map[string]interface{})["email"].(string),
		Redis: item.(map[string]interface{})["id"].(string),
	}

	// Insert the item into the collection
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		logrus.Error(err)
	}

}
