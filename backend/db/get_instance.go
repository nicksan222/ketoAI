package db

import (
	"context"
	"errors"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

// GetClient returns a mongo.Client instance, creating it on the first call.
func GetDBClient() (*mongo.Database, error) {
	var err error

	clientOnce.Do(func() {
		MONGO_CONNECTION_STRING := os.Getenv("MONGO_CONNECTION_STRING")
		if MONGO_CONNECTION_STRING == "" {
			err = errors.New("MONGO_CONNECTION_STRING not set")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientInstance, err = mongo.Connect(ctx, options.Client().ApplyURI(MONGO_CONNECTION_STRING))
		if err != nil {
			return
		}

		// Optionally, add a ping to ensure connection is alive.
		err = clientInstance.Ping(ctx, nil)
	})

	return clientInstance.Database("ketoai"), err
}
