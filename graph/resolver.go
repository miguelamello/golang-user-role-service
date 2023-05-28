package graph

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/miguelamello/user-domain-role-service/graph/model"
	"github.com/miguelamello/user-domain-role-service/graph/packages/redis"
	"github.com/miguelamello/user-domain-role-service/graph/packages/validation"
)

type Resolver struct{}

// Implement the UserById resolver function
func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {

	// Verify if id is valid
	if !validation.VerifyUUID(id) {
		return nil, errors.New("id is not valid")
	}

	// Get the user from Redis
	user, err := redis.GetUser(id)
	if err != nil {
		return nil, errors.New("user does not exist")
	}

	return user, nil

}

// Implement the CreateUser resolver function
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {

	name := strings.TrimSpace(input.Name)
	email := strings.TrimSpace(input.Email)

	// Verify if name and email are not empty
	if len(name) == 0 {
		return nil, errors.New("name is a required field")
	}

	// Verify if email are not empty
	if len(email) == 0 {
		return nil, errors.New("email is a required field")
	}

	// Verify if the email is valid
	if !validation.ValidateEmailString(email) {
		return nil, errors.New("email is not valid")
	}

	// Implement the logic to create a new user
	// Generate a unique ID for the new user
	userID := generateUniqueID()

	// Create a new user object
	newUser := &model.User{
		ID:    userID,
		Name:  name,
		Email: email,
	}

	// Save the new user to Redis
	success, err := redis.SaveUser(newUser)
	if err != nil {
		return nil, errors.New("failed to save user and already reported to team")
	}

	if !success {
		return nil, errors.New("failed to save user and already reported to team")
	}

	// Return the newly created user
	return newUser, nil

}

// Generate a unique ID for the user
func generateUniqueID() string {

	id := uuid.New()
	idString := id.String()
	return idString

}
