package main

/*

Entry point for FeedFlux.

In Phase 1, this file is responsible for:

1. Bootstrapping the HTTP server (Gin).
2. Wiring together repository implementations.
3. Injecting dependencies into services.
4. Injecting services into handlers.
5. Registering HTTP routes.
6. Starting the server.

Important architectural principle:
----------------------------------
This file performs dependency wiring only.
It does NOT contain business logic, persistence logic, or validation logic.

All layers are composed here following dependency injection,
which ensures that lower-level implementations (like in-memory storage)
can later be replaced with PostgreSQL without changing higher layers.

*/

import (
	"github.com/gin-gonic/gin"

	"github.com/Krishna8167/feedflux/internal/handler"
	"github.com/Krishna8167/feedflux/internal/repository/memory"
	"github.com/Krishna8167/feedflux/internal/service"
)


func main() {

	/*
		Gin Engine Initialization
		-------------------------
		gin.Default() initializes the HTTP router with:
		- Logging middleware
		- Recovery middleware (prevents server crash on panic)

		This represents the HTTP layer of the architecture.
	*/

	r := gin.Default()

	/*
		Repository Layer (Storage Abstraction)
		--------------------------------------
		In Phase 1, we use in-memory repositories backed by maps.

		The application depends on repository interfaces,
		not concrete implementations.

		This enables:
		- Easy swapping to PostgreSQL in Phase 2
		- Clean separation of storage from business logic
	*/

	feedRepo := memory.NewFeedMemoryRepository()
	articleRepo := memory.NewArticleMemoryRepository()

	/*
		Service Layer (Business Logic)
		-------------------------------
		Services contain application rules and orchestration logic.
		They depend only on repository interfaces.

		Responsibilities include:
		- Validation
		- Business rules
		- Coordinating persistence
	*/

	feedService := service.NewFeedService(feedRepo)
	articleService := service.NewArticleService(articleRepo)

	/*
		Handler Layer (HTTP Adapters)
		------------------------------
		Handlers translate HTTP requests into service calls.

		They are responsible for:
		- Parsing JSON input
		- Mapping errors to HTTP responses
		- Returning structured responses

		Handlers do NOT contain business logic.
	*/

	feedHandler := handler.NewFeedHandler(feedService)
	articleHandler := handler.NewArticleHandler(articleService)

	/*
		Route Registration
		------------------
		Each route maps an HTTP endpoint to a handler method.

		Current endpoints:
		- POST /feeds      -> Create a new feed
		- GET  /feeds      -> List all feeds
		- GET  /articles   -> List stored articles

		These endpoints form the API surface of FeedFlux Phase 1.
	*/

	r.POST("/feeds", feedHandler.CreateFeed)
	r.GET("/feeds", feedHandler.ListFeeds)
	r.GET("/articles", articleHandler.ListArticles)
	
	/*
		Server Startup
		--------------
		The application listens on port 8080.

		In future phases:
		- Port will be configurable via environment variables.
		- Graceful shutdown with context will be implemented.
	*/

	r.Run(":8080")
}
