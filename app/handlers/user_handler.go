package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go-restapi-unittesting/app/dto"
	"go-restapi-unittesting/app/usecases"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecaseInterface
}

func NewUserHandler(userUsecase usecases.UserUsecaseInterface) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (u UserHandler) Register(c *fiber.Ctx) error {
	//ValidateStruct
	userPayload := new(dto.CreateUserPayload)
	if err := c.BodyParser(userPayload); err != nil {
		log.Error(err)
		return errors.New("Payload Tidak Sesuai")
	}
	err := u.UserUsecase.UserRegister(c.Context(), *userPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
