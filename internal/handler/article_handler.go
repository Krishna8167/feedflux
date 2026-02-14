package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Krishna8167/feedflux/internal/service"
)

type ArticleHandler struct {
	service *service.ArticleService
}

func NewArticleHandler(s *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: s}
}

func (h *ArticleHandler) ListArticles(c *gin.Context) {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")
	feedID := c.Query("feed_id")

	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)

	var articles interface{}
	var err error

	if feedID != "" {
		articles, err = h.service.ListArticlesByFeed(feedID, limit, page)
	} else {
		articles, err = h.service.ListArticles(limit, page)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)
}
