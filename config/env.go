package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	MongoDBConnection MongoDBConnection
}

var GlobalEnv ConfigEnv

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	mongodb := MongoDBConnection{
		Url:      os.Getenv("MONGO_DB_URL"),
		Port:     os.Getenv("MONGO_DB_PORT"),
		Database: os.Getenv("MONGO_DB_DATABASE"),
	}

	GlobalEnv = ConfigEnv{
		MongoDBConnection: mongodb,
	}

}
