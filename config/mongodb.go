package config

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDBConnection struct {
	Url      string
	Port     string
	Database string
}

type GetConnection struct {
}

func NewGetConnection() *GetConnection {
	return &GetConnection{}
}

type DatabaseConnector interface {
	Connect() (*mongo.Database, error)
}

func (c *GetConnection) Connect() (*mongo.Database, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", GlobalEnv.MongoDBConnection.Url, GlobalEnv.MongoDBConnection.Port)))
	if err != nil {
		panic("Mongodb connection error")
	}
	return client.Database(GlobalEnv.MongoDBConnection.Database), nil
}
