package graph

//xxxxgo:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"errors"
	"strings"
	"github.com/google/uuid"
	"github.com/miguelamello/user-domain-role-service/graph/model"
	"github.com/miguelamello/user-domain-role-service/graph/entities"
)

type Resolver struct{}

// Implement the CreateUser resolver function
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {

	// Verify if name and email are not empty
	if strings.TrimSpace(input.Name) == "" || strings.TrimSpace(input.Email) == "" {
		return nil, errors.New("name and email are required fields")
	}

	// Verify if the email is valid
	if !validation.ValidateEmailString(input.Email) {
		return nil, errors.New("email is not valid")
	}

	// Implement the logic to create a new user
	// Generate a unique ID for the new user
	userID := generateUniqueID()

	// Create a new user object
	newUser := &model.User{
		ID:    userID,
		Name:  input.Name,
		Email: input.Email,
	}

	// Save the new user to the database or perform any other necessary actions

	// Return the newly created user
	return newUser, nil

}

// Helper function to generate a unique ID for the user
func generateUniqueID() string {

	// Generate a new UUID (version 4)
	id := uuid.New()
	// Convert the UUID to a string representation
	idString := id.String()
	return idString

}
