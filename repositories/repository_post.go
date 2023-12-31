package repositories

import (
	"context"
	"fmt"
	"github.com/tomerg2/mini-hacker-news/db_client"
	"github.com/tomerg2/mini-hacker-news/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetPosts(db *mongo.Database) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"upvotes", -1}})

	postsCollection := db.Collection(db_client.POSTS_COLLECTION)
	cursor, err := postsCollection.Find(ctx, bson.M{}, findOptions)
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
		Upvotes: 0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postsCollection := db.Collection(db_client.POSTS_COLLECTION)
	res, err := postsCollection.InsertOne(ctx, post)
	if err != nil {
		return "", err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	} else {
		return "", fmt.Errorf("failed to extract post id")
	}
}

func UpvotePost(db *mongo.Database, upvote models.Upvote) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postsCollection := db.Collection(db_client.UPVOTES_COLLECTION)
	res, err := postsCollection.InsertOne(ctx, upvote)
	if err != nil {
		return "", err
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)

	return oid.Hex(), nil
}

func CalculatePostUpvotes(db *mongo.Database, postId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postsCollection := db.Collection(db_client.POSTS_COLLECTION)
	filter := bson.M{"_id": postId}
	update := bson.M{"$inc": bson.M{"upvotes": 1}}

	_, err := postsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
