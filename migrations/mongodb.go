package migrations

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go-restapi-unittesting/config"
	"log"
)

func RunMigration() {
	m, err := migrate.New(
		"file://migrations/mongodb",
		fmt.Sprintf("mongodb://%s:%s/%s",
			config.GlobalEnv.MongoDBConnection.Url,
			config.GlobalEnv.MongoDBConnection.Port,
			config.GlobalEnv.MongoDBConnection.Database))

	if err != nil {
		panic("Migration Set Connection Error")
	}
	err = m.Up()
	if err != nil {
		log.Println(err)
	}
}
