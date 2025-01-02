package repositories

import (
	"context"
	"fmt"
	"go-restapi-unittesting/app/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

type UserRepositoryInterface interface {
	GetDataUserByName(ctx context.Context, db *mongo.Database, name string) ([]dto.Users, error)
}

func (u UserRepository) GetDataUserByName(ctx context.Context, db *mongo.Database, name string) ([]dto.Users, error) {
	var users []dto.Users

	// Define filter to find users by name
	filter := bson.M{"name": name}

	// Execute the Find operation
	cursor, err := db.Collection("users").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding users by name: %w", err)
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document into dto.Users
	for cursor.Next(ctx) {
		var user dto.Users
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("error decoding user: %w", err)
		}
		users = append(users, user)
	}

	// Check if cursor encountered any errors during iteration
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error during cursor iteration: %w", err)
	}

	return users, nil
}
