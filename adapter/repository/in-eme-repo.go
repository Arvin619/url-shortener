package repository

import (
	"context"
	"github.com/Arvin619/url-shortener/domain/entity"
	"github.com/Arvin619/url-shortener/domain/errors"
	"github.com/Arvin619/url-shortener/domain/repository"
	"sync"
)

var _ repository.IUrlShortenerRepo = (*InMemRepo)(nil)

type InMemRepo struct {
	storage map[string]*entity.UrlShortener
	rwMu    *sync.RWMutex
}

func NewInMemRepo() *InMemRepo {
	return &InMemRepo{
		storage: make(map[string]*entity.UrlShortener),
		rwMu:    &sync.RWMutex{},
	}
}

func (i *InMemRepo) Store(_ context.Context, url *entity.UrlShortener) error {
	i.rwMu.Lock()
	defer i.rwMu.Unlock()
	i.storage[url.ShortId] = url
	return nil
}

func (i *InMemRepo) Fetch(_ context.Context, shortId string) (*entity.UrlShortener, error) {
	i.rwMu.RLock()
	defer i.rwMu.RUnlock()
	u, ok := i.fetch(shortId)
	if !ok {
		return nil, errors.ErrUrlShortenerNotFound
	}
	return u, nil
}

func (i *InMemRepo) fetch(shortId string) (*entity.UrlShortener, bool) {
	u, ok := i.storage[shortId]
	return u, ok
}
