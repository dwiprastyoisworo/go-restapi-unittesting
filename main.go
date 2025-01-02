package main

import (
	config2 "go-restapi-unittesting/config"
	"go-restapi-unittesting/migrations"
	"go-restapi-unittesting/routes"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// use a single instance of Validate, it caches struct info
var ValidateStruct *validator.Validate

func main() {
	config2.LoadEnv()
	app := fiber.New()
	// get mongodb connection
	db, err := config2.NewGetConnection().Connect()
	if err != nil {
		panic(err)
	}
	// run mongodb migrations
	migrations.RunMigration()

	// user routes
	userRoute := routes.NewUserRoute(app, db)
	userRoute.AppRoute()

	// setup validator
	ValidateStruct = validator.New(validator.WithRequiredStructEnabled())

	if err := app.Listen(":8087"); err != nil {
		log.Fatal(err)
	}

}
