package entity

import "time"

// UrlShortener 短網址結構
type UrlShortener struct {
	// ShortId 短網址的 id
	ShortId string

	// OriginalUrl 原網址
	OriginalUrl string

	// CreateTime 產生時間
	CreateTime time.Time
}

func NewUrlShortener(shortId string, originalUrl string) *UrlShortener {
	return &UrlShortener{
		ShortId:     shortId,
		OriginalUrl: originalUrl,
		CreateTime:  time.Now(),
	}
}
