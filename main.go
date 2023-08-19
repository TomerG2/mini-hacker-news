package main

import (
	"github.com/tomerg2/mini-hacker-news/api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start the server")
	}
}
