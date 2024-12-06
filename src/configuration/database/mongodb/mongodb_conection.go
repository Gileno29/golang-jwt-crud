package mongodb

import (
	"context"
	"os"

	"github.com/Gileno29/golang-jwt-crud/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGODB_URL           = "MONGODB_URL"
	MONGODB_USER_DATABASE = "MONGODB_USER_DATABASE"
)

func NewMongoDBCOnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_USER_DATABASE)

	client, err := mongo.Connect(options.Client().ApplyURI(mongodb_uri))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("conseguiu conectar")

	return client.Database(mongodb_database), nil

}
