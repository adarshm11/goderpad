package db

import (
	"context"

	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"

	"goderpad/models"
)

// TODO: implement actual database logic
func GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	// query MongoDB for user by ID and return a User model
	return nil, nil
}
