package service

import (
	"errors"

	"github.com/Krishna8167/feedflux/internal/model"
	"github.com/Krishna8167/feedflux/internal/repository"
)

type ArticleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) ListArticles(limit, page int) ([]model.Article, error) {
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	return s.repo.GetAll(limit, offset)
}

func (s *ArticleService) ListArticlesByFeed(feedID string, limit, page int) ([]model.Article, error) {
	if feedID == "" {
		return nil, errors.New("feed_id is required")
	}

	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	return s.repo.GetByFeedID(feedID, limit, offset)
}
