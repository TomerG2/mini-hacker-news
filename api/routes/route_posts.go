package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tomerg2/mini-hacker-news/api/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/posts", handlers.GetPosts)
	router.POST("/posts", handlers.CreatePost)
}
