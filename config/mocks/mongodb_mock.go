package mocks

import (
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/stretchr/testify/mock"
)

// MockDatabaseConnector is a mock implementation of DatabaseConnector
type MockDatabaseConnector struct {
	mock.Mock
}

// Mock Connect method
func (m *MockDatabaseConnector) Connect() (*mongo.Database, error) {
	args := m.Called()
	if db, ok := args.Get(0).(*mongo.Database); ok {
		return db, args.Error(1)
	}
	return nil, args.Error(1)
}
