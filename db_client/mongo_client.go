package db_client

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

func GetMongoClient() (*mongo.Database, error) {
	mongoUrl := os.Getenv("MONGO_URL")
	if mongoUrl == "" {
		return nil, fmt.Errorf("MONGO_URL env var must be provided")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}

	// TODO fix close connection to mongo
	//defer func() {
	//	if err := client.Disconnect(ctx); err != nil {
	//		logrus.Error("Failed to disconnect from the database")
	//	}
	//}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client.Database(HACKER_DB), nil
}
