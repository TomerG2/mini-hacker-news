package repositories

import (
	"context"
	"fmt"
	"github.com/tomerg2/mini-hacker-news/db_client"
	"github.com/tomerg2/mini-hacker-news/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetPosts(db *mongo.Database) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postsCollection := db.Collection(db_client.POSTS_COLLECTION)
	cursor, err := postsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	posts := []models.Post{}
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func CreatePost(db *mongo.Database, content string) (string, error) {
	post := models.Post{
		Content: content,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postsCollection := db.Collection(db_client.POSTS_COLLECTION)
	res, err := postsCollection.InsertOne(ctx, post)
	if err != nil {
		return "", err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.String(), nil
	} else {
		return "", fmt.Errorf("failed to extract post id")
	}
}

func UpvotePost(db *mongo.Database) error {
	return nil
}

func CalculatePostUpvotes(db *mongo.Database, postId string) error {
	return nil
}
