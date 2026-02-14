package memory

import (
	"sync"

	"github.com/Krishna8167/feedflux/internal/model"
)

type ArticleMemoryRepository struct {
	mu       sync.RWMutex
	articles map[string]*model.Article
}

func NewArticleMemoryRepository() *ArticleMemoryRepository {
	return &ArticleMemoryRepository{
		articles: make(map[string]*model.Article),
	}
}

func (r *ArticleMemoryRepository) Create(article *model.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.articles[article.ID] = article
	return nil
}

func (r *ArticleMemoryRepository) GetAll(limit, offset int) ([]model.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []model.Article
	for _, a := range r.articles {
		list = append(list, *a)
	}

	// pagination
	end := offset + limit
	if offset > len(list) {
		return []model.Article{}, nil
	}
	if end > len(list) {
		end = len(list)
	}

	return list[offset:end], nil
}

func (r *ArticleMemoryRepository) GetByFeedID(feedID string, limit, offset int) ([]model.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var filtered []model.Article
	for _, a := range r.articles {
		if a.FeedID == feedID {
			filtered = append(filtered, *a)
		}
	}

	end := offset + limit
	if offset > len(filtered) {
		return []model.Article{}, nil
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[offset:end], nil
}
