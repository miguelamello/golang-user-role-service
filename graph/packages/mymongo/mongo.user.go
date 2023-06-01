package mymongo

import (
	"github.com/miguelamello/user-domain-role-service/graph/model"
	"github.com/sirupsen/logrus"

	//"go.mongodb.org/mongo-driver/mongo"
	"github.com/miguelamello/user-domain-role-service/graph/packages/mymongo"
)

func PushUser(user *model.User) {
	
	// Get the Redis client
	client, err := mymongo.getClient()
	if err != nil {
		logrus.Error(err)
		return false, err
	}

}