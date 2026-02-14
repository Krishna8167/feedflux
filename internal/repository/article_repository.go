package repository

import "github.com/Krishna8167/feedflux/internal/model"

type ArticleRepository interface {
	Create(article *model.Article) error
	GetAll(limit, offset int) ([]model.Article, error)
	GetByFeedID(feedID string, limit, offset int) ([]model.Article, error)
}
