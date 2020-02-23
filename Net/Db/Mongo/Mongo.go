package Mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Current instance
var client *mongo.Client
var database string

// Get current instance
func GetDatabase() *mongo.Database {
	return client.Database(database)
}

// Connection to DB
func Connect(config Config) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.ConnectionString))
	if err != nil {
		err = errors.Wrap(err, "Error connection to DB")
		return
	}

	database = config.DatabaseName

	return
}
