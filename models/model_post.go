package models

type Post struct {
	ID      string `bson:"_id" json:"_id"`
	Content string `bson:"content" json:"content"`
}
