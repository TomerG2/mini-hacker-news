package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/tomerg2/mini-hacker-news/api/dtos"
	"github.com/tomerg2/mini-hacker-news/db_client"
	"github.com/tomerg2/mini-hacker-news/models"
	"github.com/tomerg2/mini-hacker-news/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	logrus.Infof("Connecting to DB")
	dbClient, err := db_client.GetMongoClient()
	if err != nil {
		logrus.Errorf("Failed connect to DB [error=%s]", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	logrus.Infof("Fetching posts from DB")
	startTime := time.Now()
	posts, err := repositories.GetPosts(dbClient)
	if err != nil {
		logrus.Errorf("Failed to fetch posts [error=%s]", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	elapsedTime := time.Since(startTime).Milliseconds()
	logrus.Infof("Fetching posts completed [posts=%d] [milliseconds=%v]", len(posts), elapsedTime)

	response := dtos.ResponsePosts{
		Posts: posts,
	}
	c.JSON(http.StatusOK, response)
}

func CreatePost(c *gin.Context) {
	post := models.Post{}
	if err := c.BindJSON(&post); err != nil {
		logrus.Errorf("Failed extract body [error=%s]", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	logrus.Infof("Connecting to DB")
	dbClient, err := db_client.GetMongoClient()
	if err != nil {
		logrus.Errorf("Failed connect to DB [error=%s]", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	logrus.Infof("Create new post")
	startTime := time.Now()
	postId, err := repositories.CreatePost(dbClient, post.Content)
	if err != nil {
		logrus.Errorf("Failed to create post [error=%s]", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	elapsedTime := time.Since(startTime).Milliseconds()
	logrus.Infof("Create post completed [postId=%s] [milliseconds=%v]", postId, elapsedTime)

	response := dtos.ResponseCreatePost{
		ID: postId,
	}
	c.JSON(http.StatusOK, response)
}

func UpvotePost(c *gin.Context) {
	//post := models.Post{}
	//if err := c.BindJSON(&post); err != nil {
	//	logrus.Errorf("Failed extract body [error=%s]", err.Error())
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	//	return
	//}

	logrus.Infof("Connecting to DB")
	dbClient, err := db_client.GetMongoClient()
	if err != nil {
		logrus.Errorf("Failed connect to DB [error=%s]", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	logrus.Infof("Upvote post")
	startTime := time.Now()
	err = repositories.UpvotePost(dbClient)
	if err != nil {
		logrus.Errorf("Failed to Upvote post [error=%s]", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	elapsedTime := time.Since(startTime).Milliseconds()
	logrus.Infof("Upvote post completed [milliseconds=%v]", elapsedTime)

	logrus.Infof("Calculate post upvotes")
	startTime = time.Now()
	err = repositories.CalculatePostUpvotes(dbClient, "132")
	if err != nil {
		logrus.Errorf("Failed to calc post upvotes[error=%s]", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	elapsedTime = time.Since(startTime).Milliseconds()
	logrus.Infof("Calculate post upvotes completed [milliseconds=%v]", elapsedTime)

	c.JSON(http.StatusOK, gin.H{})
}
