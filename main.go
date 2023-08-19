package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tomerg2/mini-hacker-news/api/routes"
)

func main() {
	router := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logrus.Fatal("Failed to start the server")
	}
}
