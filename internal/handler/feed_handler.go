package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Krishna8167/feedflux/internal/service"

)

type FeedHandler struct {
	service *service.FeedService
}

func NewFeedHandler(s *service.FeedService) *FeedHandler {
	return &FeedHandler{service: s}
}

type createFeedRequest struct {
	URL  string `json:"url" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (h *FeedHandler) CreateFeed(c *gin.Context) {
	var req createFeedRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feed, err := h.service.AddFeed(req.URL, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, feed)
}

func (h *FeedHandler) ListFeeds(c *gin.Context) {
	feeds, err := h.service.ListFeeds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, feeds)
}
