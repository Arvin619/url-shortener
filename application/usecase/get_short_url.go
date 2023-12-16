package usecase

import (
	"context"
	"errors"
	"github.com/Arvin619/url-shortener/domain/repository"
	"regexp"
)

var shortIdRegex = regexp.MustCompile(`^[0-9a-zA-Z]{6}$`)

type GetShortUrlUseCase struct {
	repo repository.IUrlShortenerRepo
}

type GetShortUrlCommand struct {
	ShortId string
}

type GetShortUrlResponse struct {
	ShortId     string
	OriginalUrl string
}

func NewGetShortUrlUseCase(repo repository.IUrlShortenerRepo) *GetShortUrlUseCase {
	return &GetShortUrlUseCase{
		repo: repo,
	}
}

func (u *GetShortUrlUseCase) Execute(ctx context.Context, cmd *GetShortUrlCommand) (*GetShortUrlResponse, error) {
	if err := cmd.validate(); err != nil {
		return nil, err
	}
	s, err := u.repo.Fetch(ctx, cmd.ShortId)
	if err != nil {
		return nil, err
	}
	return &GetShortUrlResponse{
		ShortId:     s.ShortId,
		OriginalUrl: s.OriginalUrl,
	}, nil
}

func (cmd *GetShortUrlCommand) validate() error {
	if cmd.ShortId == "" {
		return errors.New("short id is required")
	}
	if !shortIdRegex.MatchString(cmd.ShortId) {
		return errors.New("short id is wrong format")
	}
	return nil
}
