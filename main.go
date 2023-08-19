package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define a route for the /posts endpoint
	router.GET("/posts", func(c *gin.Context) {
		// Mocking some sample data for demonstration
		posts := []map[string]interface{}{
			{"id": 1, "title": "Sample Post 1", "content": "This is the content of post 1."},
			{"id": 2, "title": "Sample Post 2", "content": "This is the content of post 2."},
		}

		c.JSON(http.StatusOK, gin.H{"posts": posts})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start the server")
	}
}
