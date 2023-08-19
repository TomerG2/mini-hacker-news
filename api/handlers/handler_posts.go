package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	// Mocking some sample data for demonstration
	posts := []map[string]interface{}{
		{"id": 1, "title": "Sample Post 1", "content": "This is the content of post 1."},
		{"id": 2, "title": "Sample Post 2", "content": "This is the content of post 2."},
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
