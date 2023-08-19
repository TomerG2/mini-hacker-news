package repositories

import (
	"context"
	"github.com/tomerg2/mini-hacker-news/models"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetPosts() ([]models.Post, error) {
	startTime := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://your_mongo_username:your_mongo_password@mongodb:27017/your_db_name"))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			logrus.Error("Failed to disconnect from the database")
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	postsCollection := client.Database("your_db_name").Collection("posts")
	cur, err := postsCollection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var posts []models.Post
	for cur.Next(ctx) {
		var post models.Post
		if err := cur.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	elapsedTime := time.Since(startTime)
	logrus.Infof("GetPosts call completed - Returned %d posts. Took %v", len(posts), elapsedTime)

	return posts, nil
}
