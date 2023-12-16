package usecase

import (
	"context"
	"errors"
	"github.com/Arvin619/url-shortener/domain/entity"
	errors2 "github.com/Arvin619/url-shortener/domain/errors"
	"github.com/Arvin619/url-shortener/domain/repository"
	"github.com/Arvin619/url-shortener/generate"
	"regexp"
)

var urlRegex = regexp.MustCompile(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`)

type AddShortUrlUseCase struct {
	repo repository.IUrlShortenerRepo
}

type AddShortUrlCommand struct {
	OriginalUrl string
}

type AddShortUrlResponse struct {
	ShortId     string
	OriginalUrl string
}

func NewAddShortUrlUseCase(repo repository.IUrlShortenerRepo) *AddShortUrlUseCase {
	return &AddShortUrlUseCase{
		repo: repo,
	}
}

func (u *AddShortUrlUseCase) Execute(ctx context.Context, cmd *AddShortUrlCommand) (*AddShortUrlResponse, error) {
	if err := cmd.validate(); err != nil {
		return nil, err
	}
	var shortId string
	for {
		shortId = generate.ShortId()
		_, err := u.repo.Fetch(ctx, shortId)
		if errors.Is(err, errors2.ErrUrlShortenerNotFound) {
			break
		}
	}
	s := entity.NewUrlShortener(shortId, cmd.OriginalUrl)
	if err := u.repo.Store(ctx, s); err != nil {
		return nil, err
	}
	return &AddShortUrlResponse{
		ShortId:     shortId,
		OriginalUrl: cmd.OriginalUrl,
	}, nil
}

func (cmd *AddShortUrlCommand) validate() error {
	if cmd.OriginalUrl == "" {
		return errors.New("original url is required")
	}
	if !urlRegex.MatchString(cmd.OriginalUrl) {
		return errors.New("original url is wrong format")
	}
	return nil
}
