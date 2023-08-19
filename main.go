package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tomerg2/mini-hacker-news/api/routes"
	"os"
)

func main() {
	router := gin.Default()
	servicePort := os.Getenv("SERVICE_PORT")

	// Initialize routes
	routes.InitializeRoutes(router)

	if err := router.Run(fmt.Sprintf(":%s", servicePort)); err != nil {
		logrus.Fatal("Failed to start the server")
	}
}
