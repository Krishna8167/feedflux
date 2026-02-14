package memory

import (
	"errors"
	"sync"

	"github.com/Krishna8167/feedflux/internal/model"
)

type FeedMemoryRepository struct {
	mu    sync.RWMutex
	feeds map[string]*model.Feed
}

func NewFeedMemoryRepository() *FeedMemoryRepository {
	return &FeedMemoryRepository{
		feeds: make(map[string]*model.Feed),
	}
}

func (r *FeedMemoryRepository) Create(feed *model.Feed) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.feeds[feed.ID] = feed
	return nil
}

func (r *FeedMemoryRepository) GetAll() ([]model.Feed, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []model.Feed
	for _, f := range r.feeds {
		result = append(result, *f)
	}
	return result, nil
}

func (r *FeedMemoryRepository) GetByID(id string) (*model.Feed, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	feed, exists := r.feeds[id]
	if !exists {
		return nil, errors.New("feed not found")
	}
	return feed, nil
}
