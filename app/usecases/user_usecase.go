package usecases

import (
	"context"
	"errors"
	"go-restapi-unittesting/app/dto"
	"go-restapi-unittesting/app/repositories"
	"go-restapi-unittesting/config"
	"go.mongodb.org/mongo-driver/v2/mongo"

	log "github.com/sirupsen/logrus"
)

type UserUsecase struct {
	db          *mongo.Database
	UserRepo    repositories.UserRepositoryInterface
	UserDtoRepo repositories.RepositoryInterface[dto.Users]
}

func NewUserUsecase(db *mongo.Database, userRepo repositories.UserRepositoryInterface, repo repositories.RepositoryInterface[dto.Users]) UserUsecaseInterface {
	return &UserUsecase{db: db, UserRepo: userRepo, UserDtoRepo: repo}
}

type UserUsecaseInterface interface {
	UserRegister(ctx context.Context, payload dto.CreateUserPayload) error
}

func (u UserUsecase) UserRegister(ctx context.Context, payload dto.CreateUserPayload) error {
	// inisiasi logger
	logger := config.NewLogger("user_usecase", "UserRegister")

	usersModel := payload.UserPayloadToUsers()

	err := u.UserDtoRepo.Create(ctx, u.db, usersModel)
	if err != nil {
		logger.SendLogger(err.Error(), log.ErrorLevel, nil)
		return errors.New("Gagal Registrasi Akun")
	}

	return nil
}
