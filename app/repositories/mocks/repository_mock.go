package mocks

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of RepositoryInterface[T]
type MockRepositoryInterface[T any] struct {
	mock.Mock
}

func (m *MockRepositoryInterface[T]) Create(ctx context.Context, db *mongo.Database, entity T) error {
	args := m.Called(ctx, db, entity)
	return args.Error(0)
}

func (m *MockRepositoryInterface[T]) Update(ctx context.Context, db *mongo.Database, filter any, update T) error {
	args := m.Called(ctx, db, filter, update)
	return args.Error(0)
}

func (m *MockRepositoryInterface[T]) Delete(ctx context.Context, db *mongo.Database, filter any) error {
	args := m.Called(ctx, db, filter)
	return args.Error(0)
}

func (m *MockRepositoryInterface[T]) GetAll(ctx context.Context, db *mongo.Database) ([]T, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]T), args.Error(1)
}

func (m *MockRepositoryInterface[T]) GetById(ctx context.Context, db *mongo.Database, id any) (T, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).(T), args.Error(1)
}
