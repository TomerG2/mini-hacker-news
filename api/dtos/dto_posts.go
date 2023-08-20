package dtos

import "github.com/tomerg2/mini-hacker-news/models"

type ResponsePosts struct {
	Posts []models.Post `bson:"posts" json:"posts"`
}

type ResponseCreatePost struct {
	ID string `json:"_id"`
}
