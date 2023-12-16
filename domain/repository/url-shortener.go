package repository

import "github.com/Arvin619/url-shortener/domain/entity"

type IUrlShortenerRepo interface {
	Store(url *entity.UrlShortener) error
	Fetch(shortId string) (*entity.UrlShortener, error)
}
