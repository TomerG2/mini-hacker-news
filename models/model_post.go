package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Content string             `bson:"content" json:"content"`
	Upvotes int                `bson:"upvotes" json:"upvotes"`
}
