package usecase

import (
	"context"
	"errors"
	"github.com/Arvin619/url-shortener/domain/repository"
	"regexp"
)

var shortIdRegex = regexp.MustCompile(`^[0-9a-zA-Z]{6}$`)

type TakeOutOriginalUrlUseCase struct {
	repo repository.IUrlShortenerRepo
}

type TakeOutOriginalUrlCommand struct {
	ShortId string
}

type TakeOutOriginalUrlResponse struct {
	ShortId     string
	OriginalUrl string
}

func NewTakeOutOriginalUrlUseCase(repo repository.IUrlShortenerRepo) *TakeOutOriginalUrlUseCase {
	return &TakeOutOriginalUrlUseCase{
		repo: repo,
	}
}

func (u *TakeOutOriginalUrlUseCase) Execute(ctx context.Context, cmd *TakeOutOriginalUrlCommand) (*TakeOutOriginalUrlResponse, error) {
	if err := cmd.validate(); err != nil {
		return nil, err
	}
	s, err := u.repo.Fetch(ctx, cmd.ShortId)
	if err != nil {
		return nil, err
	}
	return &TakeOutOriginalUrlResponse{
		ShortId:     s.ShortId,
		OriginalUrl: s.OriginalUrl,
	}, nil
}

func (cmd *TakeOutOriginalUrlCommand) validate() error {
	if cmd.ShortId == "" {
		return errors.New("short id is required")
	}
	if !shortIdRegex.MatchString(cmd.ShortId) {
		return errors.New("short id is wrong format")
	}
	return nil
}
