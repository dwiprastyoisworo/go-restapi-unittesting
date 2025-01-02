package tests

import (
	"context"
	"errors"
	"go-restapi-unittesting/app/dto"
	"go-restapi-unittesting/app/repositories/mocks"
	"go-restapi-unittesting/app/usecases"
	mocks2 "go-restapi-unittesting/config/mocks"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestUserRegister_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepositoryInterface[dto.Users])
	mockDb := &mongo.Database{}

	// Create mock database connector
	mockDbConnector := &mocks2.MockDatabaseConnector{}

	mockDbConnector.On("Connect").Return(mockDb, nil)
	db, _ := mockDbConnector.Connect()

	usecase := usecases.NewUserUsecase(db, nil, mockRepo)

	payload := dto.CreateUserPayload{
		Name:  "John Doe",
		Email: "johndoe@example.com",
		// other payload fields...
	}

	usersModel := payload.UserPayloadToUsers()
	mockRepo.On("Create", mock.Anything, mockDb, usersModel).Return(nil)

	// Act
	err := usecase.UserRegister(context.Background(), payload)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserRegister_Fail(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepositoryInterface[dto.Users])
	mockDB := &mongo.Database{} // mock DB (can be left nil for this test)

	// Create mock database connector
	mockDbConnector := &mocks2.MockDatabaseConnector{}

	mockDbConnector.On("Connect").Return(mockDB, nil)
	db, _ := mockDbConnector.Connect()

	usecase := usecases.NewUserUsecase(db, nil, mockRepo)

	payload := dto.CreateUserPayload{
		Name:  "John Doe",
		Email: "johndoe@example.com",
		// other payload fields...
	}

	usersModel := payload.UserPayloadToUsers()
	mockRepo.On("Create", mock.Anything, mockDB, usersModel).Return(errors.New("database error"))

	// Act
	err := usecase.UserRegister(context.Background(), payload)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "Gagal Registrasi Akun", err.Error())
	mockRepo.AssertExpectations(t)
}
