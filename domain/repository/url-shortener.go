package repository

import (
	"context"
	"github.com/Arvin619/url-shortener/domain/entity"
)

type IUrlShortenerRepo interface {
	Store(ctx context.Context, url *entity.UrlShortener) error
	Fetch(ctx context.Context, shortId string) (*entity.UrlShortener, error)
}
