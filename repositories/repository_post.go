package repositories

import (
	"context"
	"github.com/tomerg2/mini-hacker-news/db_client"
	"github.com/tomerg2/mini-hacker-news/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetPosts(db *mongo.Database) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postsCollection := db.Collection(db_client.POSTS_COLLECTION)
	cur, err := postsCollection.Find(ctx, bson.M{})
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

	return posts, nil
}
