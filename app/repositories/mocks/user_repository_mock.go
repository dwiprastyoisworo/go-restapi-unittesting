package mocks

import (
	"context"
	"go-restapi-unittesting/app/dto"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock for the UserRepositoryInterface
type MockUserRepository struct {
	mock.Mock
}

// GetDataUserByName mocks the GetDataUserByName method of UserRepository
func (m *MockUserRepository) GetDataUserByName(ctx context.Context, db *mongo.Database, name string) ([]dto.Users, error) {
	args := m.Called(ctx, db, name)
	return args.Get(0).([]dto.Users), args.Error(1)
}
