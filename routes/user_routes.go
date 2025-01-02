package routes

import (
	"go-restapi-unittesting/app/dto"
	"go-restapi-unittesting/app/handlers"
	"go-restapi-unittesting/app/repositories"
	"go-restapi-unittesting/app/usecases"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	App *fiber.App
	Db  *mongo.Database
}

func NewUserRoute(app *fiber.App, db *mongo.Database) *UserRoute {
	return &UserRoute{App: app, Db: db}
}

func (r UserRoute) AppRoute() {
	repository := repositories.NewRepository[dto.Users]("users")
	userRepository := repositories.NewUserRepository()
	userUsecase := usecases.NewUserUsecase(r.Db, userRepository, repository)
	userHandler := handlers.NewUserHandler(userUsecase)

	api := r.App.Group("/users")
	api.Post("/register", userHandler.Register)
}
