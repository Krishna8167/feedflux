package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Krishna8167/feedflux/internal/handler"
	"github.com/Krishna8167/feedflux/internal/repository/memory"
	"github.com/Krishna8167/feedflux/internal/service"
)


func main() {
	r := gin.Default()

	// Repositories
	feedRepo := memory.NewFeedMemoryRepository()
	articleRepo := memory.NewArticleMemoryRepository()

	// Services
	feedService := service.NewFeedService(feedRepo)
	articleService := service.NewArticleService(articleRepo)

	// Handlers
	feedHandler := handler.NewFeedHandler(feedService)
	articleHandler := handler.NewArticleHandler(articleService)

	// Routes
	r.POST("/feeds", feedHandler.CreateFeed)
	r.GET("/feeds", feedHandler.ListFeeds)
	r.GET("/articles", articleHandler.ListArticles)

	r.Run(":8080")
}
