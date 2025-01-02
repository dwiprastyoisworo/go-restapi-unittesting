package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository[T any] struct {
	CollectionName string
}

func NewRepository[T any](collectionName string) RepositoryInterface[T] {
	return &Repository[T]{CollectionName: collectionName}
}

type RepositoryInterface[T any] interface {
	Create(ctx context.Context, db *mongo.Database, entity T) error
	Update(ctx context.Context, db *mongo.Database, filter any, update T) error
	Delete(ctx context.Context, db *mongo.Database, filter any) error
	GetAll(ctx context.Context, db *mongo.Database) ([]T, error)
	GetById(ctx context.Context, db *mongo.Database, id any) (T, error)
}

// Create inserts a new document in the collection
func (r *Repository[T]) Create(ctx context.Context, db *mongo.Database, entity T) error {
	_, err := db.Collection(r.CollectionName).InsertOne(ctx, entity)
	return err
}

// Update modifies an existing document based on a filter
func (r *Repository[T]) Update(ctx context.Context, db *mongo.Database, filter any, update T) error {
	_, err := db.Collection(r.CollectionName).UpdateOne(ctx, filter, bson.M{"$set": update})
	return err
}

// Delete removes a document based on a filter
func (r *Repository[T]) Delete(ctx context.Context, db *mongo.Database, filter any) error {
	_, err := db.Collection(r.CollectionName).DeleteOne(ctx, filter)
	return err
}

// GetAll retrieves all documents in the collection as a slice of type T
func (r *Repository[T]) GetAll(ctx context.Context, db *mongo.Database) ([]T, error) {
	var results []T
	cursor, err := db.Collection(r.CollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc T
		if err := cursor.Decode(&doc); err != nil {
			return nil, fmt.Errorf("error decoding document: %w", err)
		}
		results = append(results, doc)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error during cursor iteration: %w", err)
	}

	return results, nil
}

// GetById retrieves a single document by its ID and decodes it into type T
func (r *Repository[T]) GetById(ctx context.Context, db *mongo.Database, id any) (T, error) {
	var result T
	filter := bson.M{"_id": id}
	err := db.Collection(r.CollectionName).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, fmt.Errorf("no document found with the given ID")
		}
		return result, fmt.Errorf("error finding document by ID: %w", err)
	}

	return result, nil
}
