package graph

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"github.com/miguelamello/user-domain-role-service/graph/model"
)

type Resolver struct{}

// Implement the CreateUser resolver function
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {

	// Implement the logic to create a new user
	// Generate a unique ID for the new user
	userID := generateUniqueID()

	// Create a new user object
	newUser := &model.User{
		ID:    userID,
		Name:  input.Name,
		Email: input.Email,
		// Initialize other fields as needed
	}

	// Save the new user to the database or perform any other necessary actions

	// Return the newly created user
	return newUser, nil

}

// Helper function to generate a unique ID for the user
func generateUniqueID() string {
	// Implement your logic to generate a unique ID
	return "e3ccxaw3hqo894zi1m2xb3j85w98v9dh"
}
