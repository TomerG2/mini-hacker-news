package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/tomerg2/mini-hacker-news/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	posts, err := repositories.GetPosts()
	if err != nil {
		logrus.Error("Failed to fetch posts")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	logrus.Infof("Fetch %d posts", len(posts))
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
